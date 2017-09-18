package redis

import (
	"container/list"
	"errors"
	"sync"
	"time"
)

var nowFunc = time.Now

type Pool struct {
	Dial func() (Conn, error)

	TestOnBorrow func(c Conn, t time.Time) error

	MaxIdle int

	MaxActive int

	IdleTimeout time.Duration

	Wait   bool
	mu     sync.Mutex
	cond   *sync.Cond
	closed bool
	active int

	// Stack of idleConn with most recently used at the front.
	idle list.List
}

type idleConn struct {
	c Conn
	t time.Time
}

func NewPool(newFn func() (Conn, error), maxIdle int) *Pool {
	return &Pool{Dial: newFn, MaxIdle: maxIdle}
}

func (pool *Pool) ActiveCount() int {
	pool.mu.Lock()
	active := pool.active
	pool.mu.Unlock()
	return active
}

func (p *Pool) IdleCount() int {
	p.mu.Lock()
	idle := p.idle.Len()
	p.mu.Unlock()
	return idle
}

func (p *Pool) release() {
	p.active -= 1
	if p.cond != nil {
		p.cond.Signal()
	}
}

func (p *Pool) Get() (Conn, error) {
	c, err := p.get()
	if err != nil {
		return nil, err
	}
	return &pooledConnection{p: p, c: c}, nil
}

func (p *Pool) get() (Conn, error) {
	p.mu.Lock()

	if timeout := p.IdleTimeout; timeout > 0 {
		for i, n := 0, p.idle.Len(); i < n; i++ {
			e := p.idle.Back()
			if e == nil {
				break
			}
			ic := e.Value.(idleConn)
			if ic.t.Add(timeout).After(nowFunc()) {
				break
			}
			p.idle.Remove(e)
			p.release()
			p.mu.Unlock()
			ic.c.Close()
			p.mu.Lock()
		}
	}
	for {
		for i, n := 0, p.idle.Len(); i < n; i++ {
			e := p.idle.Front()
			if e == nil {
				break
			}
			ic := e.Value.(idleConn)
			p.idle.Remove(e)
			testFunc := p.TestOnBorrow
			p.mu.Unlock()
			if testFunc == nil || testFunc(ic.c, ic.t) == nil {
				return ic.c, nil
			}
			ic.c.Close()
			p.mu.Lock()
			p.release()
		}

		if p.closed {
			p.mu.Unlock()
			return nil, errors.New("redigo: get on closed pool")
		}

		if p.MaxActive == 0 || p.active < p.MaxActive {
			dial := p.Dial
			p.active += 1
			p.mu.Unlock()
			c, err := dial()
			if err != nil {
				p.mu.Lock()
				p.release()
				p.mu.Unlock()
				c = nil
			}
			return c, err
		}

		if !p.Wait {
			p.mu.Unlock()
			return nil, errors.New("pool not conn")
		}

		if p.cond == nil {
			p.cond = sync.NewCond(&p.mu)
		}
		p.cond.Wait()
	}
}

func (p *Pool) put(c Conn) error {
	err := c.Err()
	p.mu.Lock()
	if !p.closed && err == nil {
		p.idle.PushFront(idleConn{t: nowFunc(), c: c})
		if p.idle.Len() > p.MaxIdle {
			c = p.idle.Remove(p.idle.Back()).(idleConn).c
		} else {
			c = nil
		}
	}
	if c == nil {
		if p.cond != nil {
			p.cond.Signal()
		}
		p.mu.Unlock()
		return nil
	}
	p.release()
	p.mu.Unlock()
	return c.Close()
}

type pooledConnection struct {
	p     *Pool
	c     Conn
	state int
}

func (pc *pooledConnection) Close() error {
	c := pc.c
	pc.p.put(c)
	return nil
}

func (pc *pooledConnection) Err() error {
	return pc.c.Err()
}

func (pc *pooledConnection) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	return pc.c.Do(commandName, args...)
}

func (pc *pooledConnection) Send(commandName string, args ...interface{}) error {
	return pc.c.Send(commandName, args...)
}

func (pc *pooledConnection) Flush() error {
	return pc.c.Flush()
}

func (pc *pooledConnection) Receive() (reply interface{}, err error) {
	return pc.c.Receive()
}
