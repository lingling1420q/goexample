package main

import (
	"log"
	"runtime"
	"sync"
	"time"
)

type Pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

func New(size int) (pool *Pool) {
	pool = &Pool{queue: make(chan int, size), wg: &sync.WaitGroup{}}
	return
}

func (p *Pool) Add(delta int) {
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}
	p.wg.Add(delta)
}

func (p *Pool) Done() {
	<-p.queue
	p.wg.Done()
}

func (p *Pool) Wait() {
	p.wg.Wait()
}

func main() {
	pool := New(100)
	log.Println("begin：", runtime.NumGoroutine())
	for i := 0; i < 1000; i++ {
		pool.Add(1)
		go func() {
			time.Sleep(time.Second)
			log.Println("running:", runtime.NumGoroutine())
			pool.Done()
		}()
	}
	pool.Wait()
	log.Println("end:", runtime.NumGoroutine())
}
