package main

import (
	"./httputil"
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

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

func handleErr(err error) {
	fmt.Println(err)
	//os.Exit(1)
}

func handleConnection(conn net.Conn) {
	var buf bytes.Buffer
	buffer := make([]byte, 8192)
	for {
		sizenew, err := conn.Read(buffer)
		buf.Write(buffer[:sizenew])
		if err == io.EOF || sizenew < 8192 {
			break
		}
	}

	content := fmt.Sprintf("%#v", buf.String())
	//fmt.Printf("content %s:", content)
	req := httputil.ParseRequest(content)
	header := make(http.Header)
	body := req.PostData
	if req.Method == "GET" {
		body = fmt.Sprintf("data %s", time.Now().Format(time.RFC1123))
	}
	headerStr := httputil.WriteHeader([]string{"", "200", "msg"}, header, body)
	//fmt.Printf("headerStr %#v", headerStr)
	conn.Write([]byte(headerStr + body))
	//conn.Write([]byte(req.PostData))
	conn.Close()
}
