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
	//"strings"
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
	request         *HTTPRequest
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
	//logs.Log.Debug("readToBuffer %#v", string(self.readBuffer.Bytes()))
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

func (self *HttpConnect) readBytes(loc int64, callback func(content []byte)) {
	self.readToBuffer()
	if self.readBufferSize > 0 {
		content := self.consume(loc)
		callback(content)
	}
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
		httpconnect := &HttpConnect{Connect: conn, Callback: self.Callback}
		httpconnect.Run()
	}
	// ch := make(chan int)
	// <-ch
}

// 2017/08/31 19:31:24 [httpserver.go:79] [D] readToBuffer "POST / HTTP/1.1\r\nHost: 120.26.13.218:8888\r\nConnection: keep-alive\r\nContent-Length: 140\r\nPostman-Token: 86025dae-725f-8602-aaca-a81fc14ed3f3\r\nCache-Control: no-cache\r\nOrigin: chrome-extension://fhbjgbiflinjbdggehcddcbncdddomop\r\nUser-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.101 Safari/537.36\r\nContent-Type: multipart/form-data; boundary=----WebKitFormBoundaryL2Qz1C5HSjjrAd1c\r\nAccept: */*\r\nAccept-Encoding: gzip, deflate\r\nAccept-Language: en-US,en;q=0.8,zh-CN;q=0.6,zh;q=0.4,vi;q=0.2,zh-TW;q=0.2\r\n\r\n------WebKitFormBoundaryL2Qz1C5HSjjrAd1c\r\nContent-Disposition: form-data; name=\"test\"\r\n\r\nasdfa\r\n------WebKitFormBoundaryL2Qz1C5HSjjrAd1c--\r\n"
// 2017/08/31 19:31:24 [httpserver.go:112] [D] consume "------WebKitFormBoundaryL2Qz1C5HSjjrAd1c\r\nContent-Disposition: form-data; name=\"test\"\r\n\r\nasdfa\r\n------WebKitFormBoundaryL2Qz1C5HSjjrAd1c--\r\n"
// 2017/08/31 19:31:24 [httpserver.go:113] [D] loc "------WebKitFormBoundaryL2Qz1C5HSjjrAd1c\r\nContent-Disposition: form-data; name=\"test\"\r\n\r\nasdfa\r\n------WebKitFormBoundaryL2Qz1C5HSjjrAd1c--\r\n\nCache-Control: no-cache\r\nOrigin: chrome-extension://fhbjgbiflinjbdggehcddcbncdddomop\r\nUser-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.101 Safari/537.36\r\nContent-Type: multipart/form-data; boundary=----WebKitFormBoundaryL2Qz1C5HSjjrAd1c\r\nAccept: */*\r\nAccept-Encoding: gzip, deflate\r\nAccept-Language: en-US,en;q=0.8,zh-CN;q=0.6,zh;q=0.4,vi;q=0.2,zh-TW;q=0.2\r\n\r\n"
