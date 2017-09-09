package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

type Conn struct {

	// Shared
	mu      sync.Mutex
	pending int
	err     error
	conn    net.Conn

	// Read
	readTimeout time.Duration
	br          *bufio.Reader

	// Write
	writeTimeout time.Duration
	bw           *bufio.Writer

	// Scratch space for formatting argument length.
	// '*' or '$', length, "\r\n"
	lenScratch [32]byte

	// Scratch space for formatting integers and floats.
	numScratch [40]byte
}

func NewConn(network, address string, options map[string]string) (*Conn, error) {
	netConn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	c := &Conn{
		conn: netConn,
		bw:   bufio.NewWriter(netConn),
		br:   bufio.NewReader(netConn),
	}
	return c, nil
}

func (c *Conn) writeLen(prefix byte, n int) error {
	c.lenScratch[len(c.lenScratch)-1] = '\n'
	c.lenScratch[len(c.lenScratch)-2] = '\r'
	i := len(c.lenScratch) - 3
	for {
		c.lenScratch[i] = byte('0' + n%10)
		i -= 1
		n = n / 10
		if n == 0 {
			break
		}
	}
	c.lenScratch[i] = prefix
	_, err := c.bw.Write(c.lenScratch[i:])
	return err
}

func (c *Conn) writeString(s string) error {
	c.writeLen('$', len(s))
	c.bw.WriteString(s)
	_, err := c.bw.WriteString("\r\n")
	return err
}

func (c *Conn) writeBytes(p []byte) error {
	c.writeLen('$', len(p))
	c.bw.Write(p)
	_, err := c.bw.WriteString("\r\n")
	return err
}

func (c *Conn) writeInt64(n int64) error {
	return c.writeBytes(strconv.AppendInt(c.numScratch[:0], n, 10))
}

func (c *Conn) writeFloat64(n float64) error {
	return c.writeBytes(strconv.AppendFloat(c.numScratch[:0], n, 'g', -1, 64))
}

func (c *Conn) fatal(err error) error {
	c.mu.Lock()
	if c.err == nil {
		c.err = err
		// Close connection to force errors on subsequent calls and to unblock
		// other reader or writer.
		c.conn.Close()
	}
	c.mu.Unlock()
	return err
}

func (c *Conn) writeCommand(cmd string, args ...interface{}) error {
	c.writeLen('$', len(cmd))
	err := c.writeString(cmd)
	for _, arg := range args {
		if err != nil {
			break
		}
		switch arg := arg.(type) {
		case string:
			err = c.writeString(arg)
		case []byte:
			err = c.writeBytes(arg)
		case int:
			err = c.writeInt64(int64(arg))
		case int64:
			err = c.writeInt64(arg)
		case float64:
			err = c.writeFloat64(arg)
		case bool:
			if arg {
				err = c.writeString("1")
			} else {
				err = c.writeString("0")
			}
		case nil:
			err = c.writeString("")
		default:
			var buf bytes.Buffer
			fmt.Fprint(&buf, arg)
			err = c.writeBytes(buf.Bytes())
		}
	}
	return err
}

func (c *Conn) Do(cmd string, args ...interface{}) (interface{}, error) {
	c.mu.Lock()
	//pending := c.pending
	c.pending = 0
	c.mu.Unlock()

	if cmd != "" {
		if err := c.writeCommand(cmd, args); err != nil {
			return nil, c.fatal(err)
		}
	}
	if err := c.bw.Flush(); err != nil {
		return nil, c.fatal(err)
	}

	var err error
	var reply interface{}
	return reply, err
}

func writeLen(prefix byte, n int) []byte {
	lenScratch := [32]byte{}
	lenScratch[len(lenScratch)-1] = '\n'
	lenScratch[len(lenScratch)-2] = '\r'
	i := len(lenScratch) - 3
	for {
		lenScratch[i] = byte('0' + n%10)
		i -= 1
		n = n / 10
		if n == 0 {
			break
		}
	}
	lenScratch[i] = prefix

	return lenScratch[i:]
}

func send(c net.Conn, cmd string) string {
	var result []byte
	c.Write([]byte(cmd))
	buffer := make([]byte, 1024*8)
	n, _ := c.Read(buffer)
	fmt.Printf("%#v\n", string(buffer[:n]))
	return string(result)
}

func writeString(s string) error {
	c.writeLen('$', len(s))
	c.bw.WriteString(s)
	_, err := c.bw.WriteString("\r\n")
	return err
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6379")
	fmt.Println(conn, err)
	//fmt.Println(string(writeLen('*', 1+len("waqu@test"))))
	// send(conn, "auth waqu@test\r\n")
	// send(conn, "select 8\r\n")
	// send(conn, "keys *\r\n")

}
