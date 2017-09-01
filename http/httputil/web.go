package httputil

import (
	"bytes"
	"encoding/json"
	"fmt"
	logs "github.com/yangaowei/gologs"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	SUPPORTED_METHODS = []string{"GET", "HEAD", "POST", "DELETE", "PATCH", "PUT", "OPTIONS"}
)

type RequestHandler struct {
	request        *HTTPRequest
	headersWritten bool
	header         http.Header
	finish         bool
	statusCode     int
	writeBuffer    bytes.Buffer
	Methods        []string
}

func NewRequestHandler(req *HTTPRequest, methods []string) *RequestHandler {
	if methods == nil {
		methods = SUPPORTED_METHODS
	}
	handler := &RequestHandler{request: req, Methods: methods, statusCode: 200}
	handler.header = make(http.Header)
	handler.header.Add("Server", "golang/server")
	handler.header.Add("Date", time.Now().Format(time.RFC1123))
	handler.header.Add("Content-Type", "text/html; charset=UTF-8")
	if value, ok := handler.request.Headers["Connection"]; ok && strings.ToLower(value[0]) == "keep-alive" {
		handler.header.Add("Connection", "Keep-Alive")
	}
	return handler
}

func (self *RequestHandler) flush() {
	var headers bytes.Buffer
	if !self.headersWritten {
		self.headersWritten = true
		header := self.generateHeaders()
		headers.Write(header)
	}
	self.request.Connection.Connect.Write(headers.Bytes())
	self.request.Connection.Connect.Write(self.writeBuffer.Bytes())
	self.request.Connection.Connect.Close()
}

func (self *RequestHandler) generateHeaders() []byte {
	var headers bytes.Buffer
	startLine := fmt.Sprintf("%s %d %s%s", self.request.Version, self.statusCode, "", CRLF)
	headers.Write([]byte(startLine))
	for key, value := range self.header {
		//fmt.Println(key, value)
		line := fmt.Sprintf("%s: %s%s", key, value[0], CRLF)
		headers.Write([]byte(line))
	}
	headers.Write([]byte(CRLF))
	return headers.Bytes()
}

func (self *RequestHandler) Finish(content interface{}) {
	var result []byte
	if reflect.ValueOf(content).Kind().String() == "map" {
		self.header.Set("Content-Type", "application/json; charset=UTF-8")
		result, _ = json.Marshal(content)
	} else {
		result = []byte(content.(string))
	}
	// if self.request.Body != nil {
	// 	result = append(result, self.request.Body.Bytes()...)
	// }
	self.Write(result)
	if _, ok := self.header["Content-Length"]; !ok {
		self.header.Add("Content-Length", strconv.Itoa(len(result)))
	}
	self.flush()
	logs.Log.Debug("finish %v", string(result))
}

func (self *RequestHandler) Write(content []byte) {
	self.writeBuffer.Write(content)
}

//POST / HTTP/1.1\r\nHost: 120.26.13.218:8888\r\nConnection: keep-alive\r\nContent-Length: 140\r\nPostman-Token: 4a3ac92c-a74e-5921-d97b-2e5fb831a3ca\r\nCache-Control: no-cache\r\nOrigin: chrome-extension://fhbjgbiflinjbdggehcddcbncdddomop\r\nUser-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.101 Safari/537.36\r\nContent-Type: multipart/form-data; boundary=----WebKitFormBoundarycc9rXTRINg9iTFHz\r\nAccept: */*\r\nAccept-Encoding: gzip, deflate\r\nAccept-Language: en-US,en;q=0.8,zh-CN;q=0.6,zh;q=0.4,vi;q=0.2,zh-TW;q=0.2\r\n\r\n------WebKitFormBoundarycc9rXTRINg9iTFHz\r\nContent-Disposition: form-data; name=\"test\"\r\n\r\nasdfa\r\n------WebKitFormBoundarycc9rXTRINg9iTFHz--\r\n
