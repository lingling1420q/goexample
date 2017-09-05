package httputil

import (
	"bytes"
	"fmt"
	logs "github.com/yangaowei/gologs"
	"io"
	"net"
	//"net/http"
	// "net/url"
	"os"
	"strconv"
	"time"
	//"strings"
)

type Server interface {
}

var (
	ListenSize    = 128
	ReadChunkSize = 1024 * 32
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
	request         *HTTPRequest
}

func HandlerError(e error) {
	logs.Log.Debug("error %v", e)
	os.Exit(1)
}

func NewHTTPConnect(conn *net.TCPConn, callback func(req *HTTPRequest)) *HttpConnect {
	httpConnect := &HttpConnect{Connect: conn, Callback: callback}
	httpConnect.Connect.SetKeepAlive(true)
	httpConnect.Connect.SetKeepAlivePeriod(3 * time.Minute)
	return httpConnect
}

func (self *HttpConnect) Run() {
	self.readToBuffer()
	self.readUntil("\r\n\r\n", self.headerCallback)
}

func (self *HttpConnect) headerCallback(headerByte []byte) {
	//logs.Log.Debug("content %#v", string(headerByte))
	startLinePoint := bytes.Index(headerByte, []byte(CRLF))
	startLine := headerByte[:startLinePoint]
	v := bytes.Split(startLine, []byte(" "))
	method, path, version := string(v[0]), string(v[1]), string(v[2])
	logs.Log.Debug("%s,%s,%s", method, path, version)
	endLinePoint := bytes.Index(headerByte, []byte(CRLF+CRLF))
	header := parseHeader(headerByte[startLinePoint+len([]byte(CRLF)) : endLinePoint])
	req := NewHTTPRequest(self, method, path, version, self.Connect.RemoteAddr().String(), header)
	self.request = req
	if c, ok := header["Content-Length"]; ok {
		contentLength, _ := strconv.ParseInt(c[0], 10, 64)
		fmt.Println("Content-Length:", contentLength)
		self.readBytes(contentLength, self.bodyCallback)
	}
	self.Callback(self.request)
}

func (self *HttpConnect) bodyCallback(bodyByte []byte) {
	var contentType string
	//logs.Log.Debug("bodyCallback %#v", string(bodyByte))
	if ct, ok := self.request.Headers["Content-Type"]; ok {
		contentType = ct[0]
		args, files := parseBody(contentType, bodyByte)
		for key, value := range args {
			self.request.Arguments.Add(key, value[0])
		}
		self.request.Files = files
		self.request.Body = bytes.NewBuffer(bodyByte)
	}
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
	// logs.Log.Debug("readToBuffer len %d", len(self.readBuffer.Bytes()))
	// logs.Log.Debug("readBufferSize len %d", self.readBufferSize)
	return self.readBufferSize
}

func (self *HttpConnect) readFromFd() (chunk []byte) {
	if self.closed {
		return
	}
	buffer := make([]byte, ReadChunkSize)
	//fmt.Println("start read .....")
	sizenew, err := self.Connect.Read(buffer)
	//fmt.Println("sizenew", sizenew)
	if err == io.EOF || sizenew < ReadChunkSize {
		self.closed = true
	}
	if sizenew > 0 {
		chunk = buffer[:sizenew]
	}
	if sizenew == 0 {
		self.closed = true
	}
	//logs.Log.Debug("buffer %s", string(buffer[:sizenew]))
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

func (self *HttpConnect) readBytes(loc int64, callback func(content []byte)) {
	fmt.Println(self.readBufferSize)
	for {
		if loc > int64(len(self.readBuffer.Bytes())) {
			self.closed = false
			self.readToBuffer()
		} else {
			break
		}
	}
	content := self.consume(loc)
	callback(content)
}

func (self *HttpConnect) consume(loc int64) (content []byte) {
	var n bytes.Buffer
	c := self.readBuffer.Bytes()
	n.Write(c[loc:])
	self.readBufferSize -= loc
	self.readBuffer = n
	//fmt.Println("readBufferSize:", self.readBufferSize)
	return c[:loc]
}

func (self *HttpConnect) Write(chunk []byte) {
	self.writeBuffer.Write(chunk)
}

func (self *HttpConnect) Finish() {
	self.Connect.Write(self.writeBuffer.Bytes())
	self.Connect.Close()
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
		httpconnect := NewHTTPConnect(conn, self.Callback)
		httpconnect.Run()
	}
	// ch := make(chan int)
	// <-ch
}
