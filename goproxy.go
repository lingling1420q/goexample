package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

var (
	hosts = "182.92.110.245:22"
)

func handleErr(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func handleConnection(conn net.Conn) {

	//dst, err := net.Dial("tcp", hosts)
	dst, err := net.DialTimeout("tcp", hosts, time.Second*10)
	if err != nil {
		handleErr(err)
	}
	go io.Copy(dst, conn)
	io.Copy(conn, dst)
	defer func() {
		conn.Close()
		dst.Close()
		fmt.Println("close conn")
	}()
}

func main() {
	ln, err := net.Listen("tcp", ":801")
	fmt.Println("listen on 801")
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
