package main

import (
	"log"
	"net"
	"time"
)

func server() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Println("error listen:", err)
		return
	}
	defer l.Close()
	log.Println("listen ok")

	for {
		time.Sleep(time.Second * 10000)
		var c net.Conn
		if c, err = l.Accept(); err != nil {
			log.Println("accept error:", err)
			break
		}
		defer c.Close()
		for {
			// read from the connection
			time.Sleep(5 * time.Second)
			var buf = make([]byte, 60000)
			log.Println("start to read from conn")
			n, err := c.Read(buf)
			if err != nil {
				log.Printf("conn read %d bytes,  error: %s", n, err)
				if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
					continue
				}
			}

			log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
		}
	}
}

func establishConn(i int) net.Conn {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Printf("%d: dial error: %s", i, err)
		return nil
	}
	log.Println(i, ":connect to server ok")
	return conn
}

func client() {
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	defer conn.Close()
	log.Println("dial ok")

	data := make([]byte, 65536)
	var total int
	for {
		n, err := conn.Write(data)
		if err != nil {
			total += n
			log.Printf("write %d bytes, error:%s\n", n, err)
			break
		}
		total += n
		log.Printf("write %d bytes this time, %d bytes in total\n", n, total)
	}

	log.Printf("write %d bytes in total\n", total)
	time.Sleep(time.Second * 10000)
}

func main() {
	go server()
	client()
}
