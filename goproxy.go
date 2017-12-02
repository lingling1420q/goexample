package main

import (
	"fmt"
	"io"
	"net"
	//"os"
	"time"
)

var (
	hosts = "47.88.100.3:8388"
)

func handleErr(err error) {
	fmt.Println(err)
	//os.Exit(1)
}

func handleConnection(conn net.Conn) {

	//dst, err := net.Dial("tcp", hosts)
	dst, err := net.DialTimeout("tcp", hosts, time.Second*1000)
	if err != nil {
		handleErr(err)
	}
	defer func() {
		conn.Close()
		dst.Close()
		fmt.Println("close conn")
	}()
	go io.Copy(dst, conn)
	io.Copy(conn, dst)

}

func main() {
	ln, err := net.Listen("tcp", ":8388")
	fmt.Println("listen on 8388")
	if err != nil {
		handleErr(err)
	}
	for {
		conn, err := ln.Accept()
		fmt.Printf("accept conn %v\n", conn)
		if err != nil {
			handleErr(err)
		}
		go handleConnection(conn)
	}
}
