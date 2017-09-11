package tcp

import (
	"fmt"
	"net"
	"time"
)

func handlerErr(err error) {
	if err != nil {
		fmt.Printf("error %v", err)
	}
}

func clinet() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:9001")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	conn.SetNoDelay(true)
	handlerErr(err)
	time.Sleep(1 * time.Second)
	fmt.Println("clinet begin write")
	conn.Write([]byte("hello world!!!"))
	conn.Close()
	fmt.Println("client close")
}
