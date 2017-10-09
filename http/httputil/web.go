package httputil

import (
	"bytes"
	"encoding/json"
	"fmt"
	logs "github.com/yangaowei/gologs"
	"net/http"
	"reflect"
	"strconv"
	//"strings"
	"sync"
	"time"
)

var (
	SUPPORTED_METHODS = []string{"GET", "HEAD", "POST", "DELETE", "PATCH", "PUT", "OPTIONS"}
)

type Handler interface {
	Finish(rh *RequestHandler)
	String() string
}

type RequestHandler struct {
	Request        *HTTPRequest
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
	handler := &RequestHandler{request: req, Request: req, Methods: methods, statusCode: 200}
	handler.header = make(http.Header)
	handler.header.Add("Server", "golang/server")
	handler.header.Add("Date", time.Now().Format(time.RFC1123))
	handler.header.Add("Content-Type", "text/html; charset=UTF-8")
	//if value, ok := handler.request.Headers["Connection"]; ok && strings.ToLower(value[0]) == "keep-alive" {
	handler.header.Add("Connection", "keep-alive")
	//}
	return handler
}

func (self *RequestHandler) flush() {
	defer self.request.Connection.Connect.Close()
	var headers bytes.Buffer
	if !self.headersWritten {
		self.headersWritten = true
		header := self.generateHeaders()
		headers.Write(header)
	}
	self.request.Connection.Connect.Write(headers.Bytes())
	self.request.Connection.Connect.Write(self.writeBuffer.Bytes())

}

func (self *RequestHandler) String() string {
	return fmt.Sprintf("RequestHandler")
}

func (self *RequestHandler) generateHeaders() []byte {
	var headers bytes.Buffer
	startLine := fmt.Sprintf("%s %d %sOK%s", self.request.Version, self.statusCode, "", CRLF)
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
	//self.Request.Connection.Run()
}

func (self *RequestHandler) Write(content []byte) {
	self.writeBuffer.Write(content)
}

type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry
	hosts bool // whether any patterns contain hostnames
}

type muxEntry struct {
	explicit bool
	h        Handler
	pattern  string
}

func (self *ServeMux) Handle(pattern string, handler Handler) {
	self.mu.Lock()
	defer self.mu.Unlock()
	self.m[pattern] = muxEntry{explicit: true, h: handler, pattern: pattern}
}

func (mux *ServeMux) match(path string) (h Handler, pattern string) {
	// Check for exact match first.
	v, ok := mux.m[path]
	if ok {
		return v.h, v.pattern
	}

	// Check for longest valid match.
	var n = 0
	for k, v := range mux.m {
		if !pathMatch(k, path) {
			continue
		}
		if h == nil || len(k) > n {
			n = len(k)
			h = v.h
			pattern = v.pattern
			fmt.Println(h, pattern)
		}
	}
	return
}

func (mux *ServeMux) execute(req *HTTPRequest) {
	path := req.Path
	handler, _ := mux.match(path)
	if handler != nil {
		handler.Finish(NewRequestHandler(req, nil))
	} else {
		h := NewRequestHandler(req, nil)
		h.statusCode = 404
		h.Finish("404")
	}
	// handler.Finish(map[string]interface{}{"result": "成功!!!!", "data": result, "files": handler.GetBase().Request.Files})
}

func pathMatch(pattern, path string) bool {
	if len(pattern) == 0 {
		// should not happen
		return false
	}
	n := len(pattern)
	//fmt.Println("pattern:", pattern, path)
	if pattern[n-1] != '/' {
		return pattern == path
	}
	return len(path) >= n && path[0:n] == pattern
}

type HandlerFunc func(content interface{})

// ServeHTTP calls f(w, r).
func (f HandlerFunc) Finish(content interface{}) {
	f(content)
}

var DefaultServeMux = &defaultServeMux

var defaultServeMux ServeMux

func NewServeMux() {
	defaultServeMux = ServeMux{m: make(map[string]muxEntry)}
}

func HandleFunc(pattern string, handler Handler) {
	DefaultServeMux.Handle(pattern, handler)
}

func Application(req *HTTPRequest) {
	//fmt.Println("Application", req)
	//fmt.Println("m", DefaultServeMux.m)
	DefaultServeMux.execute(req)
	// handler := NewRequestHandler(req, nil)
	// //fmt.Println(handler)
	// //handler.Finish(map[string]string{"msg": "hellow world!"})
	// result := handler.Request.Arguments
	// handler.Finish(map[string]interface{}{"result": "成功", "data": result, "files": handler.Request.Files})
}
