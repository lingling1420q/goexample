package httputil

import (
	"bytes"
	"fmt"
	logs "github.com/yangaowei/gologs"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

type Server interface {
}

var (
	ListenSize    = 128
	ReadChunkSize = 8192
	//CRLF          = "\r\n"
	//RequestCallback func(r *HTTPRequest)
)

type HttpServer struct {
	Port      int
	Host      string
	KeepAlive bool
	Protocol  string
	Callback  func(req *HTTPRequest)
}

type HttpConnect struct {
	Connect         *net.TCPConn
	Address         string
	Protocol        string
	readBuffer      bytes.Buffer
	readBufferSize  int64
	writeBuffer     bytes.Buffer
	writeBufferSize int64
	closed          bool
	Callback        func(req *HTTPRequest)
}

func HandlerError(e error) {
	logs.Log.Debug("error %v", e)
	os.Exit(1)
}

func (self *HttpConnect) Run() {
	self.readToBuffer()
	self.readUntil("\r\n\r\n", self.headerCallback)
}

func (self *HttpConnect) headerCallback(headerByte []byte) {
	startLinePoint := bytes.Index(headerByte, []byte(CRLF))
	startLine := headerByte[:startLinePoint]
	v := bytes.Split(startLine, []byte(" "))
	method, path, version := string(v[0]), string(v[1]), string(v[2])
	logs.Log.Debug("%s,%s,%s", method, path, version)
	endLinePoint := bytes.Index(headerByte, []byte(CRLF+CRLF))
	header := make(http.Header)
	headerByte = headerByte[startLinePoint+len([]byte(CRLF)) : endLinePoint]
	for _, line := range bytes.Split(headerByte, []byte(CRLF)) {
		l := bytes.SplitN(line, []byte(":"), 2)
		header.Add(string(l[0]), strings.TrimLeft(string(l[1]), " "))
	}
	logs.Log.Debug("header %#v", header)
	//req := &HTTPRequest{Connection: self, Method: method, Uri: path, Version: version, Headers: header, RemoteIp: self.Connect.RemoteAddr().String()}
	req := NewHTTPRequest(self, method, path, version, self.Connect.RemoteAddr().String(), header)
	self.Callback(req)
}

func (self *HttpConnect) readToBuffer() (size int64) {
	for {
		chuck := self.readFromFd()
		if len(chuck) == 0 {
			break
		}
		self.readBuffer.Write(chuck)
		self.readBufferSize += int64(len(chuck))
	}
	return self.readBufferSize
}

func (self *HttpConnect) readFromFd() (chunk []byte) {
	if self.closed {
		return
	}
	buffer := make([]byte, ReadChunkSize)
	sizenew, err := self.Connect.Read(buffer)
	if err == io.EOF || sizenew < ReadChunkSize {
		self.closed = true
	}
	chunk = buffer[:sizenew]
	return
}

func (self *HttpConnect) readUntil(delimiter string, callback func(content []byte)) {
	self.readToBuffer()
	if self.readBufferSize > 0 {
		del := []byte(delimiter)
		point := bytes.Index(self.readBuffer.Bytes(), del)
		loc := point + len(del)
		content := self.consume(int64(loc))
		callback(content)
	}
}

func (self *HttpConnect) consume(loc int64) (content []byte) {
	c := self.readBuffer.Bytes()
	self.readBuffer.Reset()
	self.readBuffer.Write(c[loc:])
	self.readBufferSize -= loc
	return c[:loc]
}

func (self *HttpServer) Listen() {
	listen := fmt.Sprintf("%s:%d", self.Host, self.Port)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", listen)
	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		HandlerError(err)
	}
	logs.Log.Debug("start server on %s", listen)
	for {
		conn, _ := ln.AcceptTCP()
		httpconnect := &HttpConnect{Connect: conn, Callback: self.Callback}
		httpconnect.Run()
	}
	// ch := make(chan int)
	// <-ch
}
