package tcp

import (
	"fmt"
	"net"
	"time"
)

var (
	Host = "0.0.0.0"
	Port = 9001
)

func handler(conn *net.TCPConn) {
	defer conn.Close()
	//conn.SetNoDelay(true)
	conn.SetNoDelay(true)
	conn.Close()
	fmt.Println("conn close")
	time.Sleep(1 * time.Second)
	for {
		fmt.Println("server begin read")
		buffer := make([]byte, 1024)
		size, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("read error:", err)
			break
		} else {
			fmt.Println("read data:", string(buffer[:size]))
			n, e := conn.Write([]byte("back!!!"))
			handlerErr(e)
			fmt.Println(n)
		}
	}
}

func server() {
	listen := fmt.Sprintf("%s:%d", Host, Port)
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", listen)
	server, err := net.ListenTCP("tcp", tcpAddr)
	fmt.Println("listen:", listen)
	if err == nil {
		for {
			conn, _ := server.AcceptTCP()
			handler(conn)
		}
	} else {
		fmt.Println("error: ", err)
	}
}
