package httputil

import (
	"bytes"
	"encoding/json"
	"fmt"
	logs "github.com/yangaowei/gologs"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

// connection=self, method=method, uri=uri, version=version,
//                 headers=headers, remote_ip=remote_ip, protocol=self.protocol
type HTTPRequest struct {
	Connection     *HttpConnect
	Method         string
	Uri            string
	Version        string
	Headers        http.Header
	RemoteIp       string
	protocol       string
	headersWritten bool
	writeBuffer    bytes.Buffer
	ResHeaders     http.Header
}

func NewHTTPRequest(args ...interface{}) *HTTPRequest {
	req := &HTTPRequest{Connection: args[0].(*HttpConnect), Method: args[1].(string), Uri: args[2].(string), Version: args[3].(string), RemoteIp: args[4].(string), Headers: args[5].(http.Header)}
	req.ResHeaders = make(http.Header)
	req.ResHeaders.Add("Server", "golang/server")
	req.ResHeaders.Add("Date", time.Now().Format(time.RFC1123))
	if _, ok := req.Headers["Content-Typ"]; !ok {
		req.ResHeaders.Add("Content-Type", "text/html; charset=UTF-8")
	}
	return req
}

func (self *HTTPRequest) String() string {
	return fmt.Sprintf("%s,%s,%s remoteIp:%s", self.Method, self.Uri, self.Version, self.RemoteIp)
}

// "Server": "TornadoServer/%s" % tornado.version,
// "Content-Type": "text/html; charset=UTF-8",
// "Date": httputil.format_timestamp(time.time()),

func (self *HTTPRequest) flush() {
	var headers bytes.Buffer
	if !self.headersWritten {
		self.headersWritten = true
		header := self.generateHeaders()
		headers.Write(header)
	}
	self.Connection.Connect.Write(headers.Bytes())
	self.Connection.Connect.Write(self.writeBuffer.Bytes())
}

func (self *HTTPRequest) generateHeaders() []byte {
	var headers bytes.Buffer
	startLine := fmt.Sprintf("%s %d %s", self.Version, 200, "")
	headers.Write([]byte(startLine))
	for key, value := range self.ResHeaders {
		line := fmt.Sprintf("%s: %s%s", key, value[0], CRLF)
		headers.Write([]byte(line))
	}
	headers.Write([]byte(CRLF))
	return headers.Bytes()
}

func (self *HTTPRequest) Finish(content interface{}) {
	if reflect.ValueOf(content).Kind().String() == "map" {
		self.ResHeaders.Set("Content-Type", "application/json; charset=UTF-8")
	}
	b, _ := json.Marshal(content)
	self.Write(b)
	if _, ok := self.ResHeaders["Content-Length"]; !ok {
		self.ResHeaders.Add("Content-Length", strconv.Itoa(len(b)))
	}
	self.flush()
	logs.Log.Debug("finish %v", string(b))
}

func (self *HTTPRequest) Write(content []byte) {
	self.writeBuffer.Write(content)
}
