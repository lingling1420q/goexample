package httputil

import (
	//"encoding/json"
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	//logs "github.com/yangaowei/gologs"
)

// connection=self, method=method, uri=uri, version=version,
//                 headers=headers, remote_ip=remote_ip, protocol=self.protocol
type HTTPRequest struct {
	Connection *HttpConnect
	Method     string
	Uri        string
	Path       string
	Version    string
	Headers    http.Header
	RemoteIp   string
	protocol   string
	Body       *bytes.Buffer
	Arguments  url.Values
	Files      map[string]interface{}
}

func NewHTTPRequest(args ...interface{}) *HTTPRequest {
	req := &HTTPRequest{Connection: args[0].(*HttpConnect), Method: args[1].(string), Uri: args[2].(string), Version: args[3].(string), RemoteIp: args[4].(string), Headers: args[5].(http.Header)}
	req.Arguments = make(url.Values)
	qs := parseQuery(req.Uri)
	if qs != nil {
		req.Arguments = qs
	}
	req.Path = strings.SplitN(req.Uri, "?", 2)[0]
	return req
}

func (self *HTTPRequest) String() string {
	return fmt.Sprintf("%s,%s,%s remoteIp:%s", self.Method, self.Path, self.Version, self.RemoteIp)
}

func (self *HTTPRequest) Finish(content []byte) {
	//b, _ := json.Marshal(content)
	self.Write(content)
	self.Connection.Finish()
}

func (self *HTTPRequest) Write(content []byte) {
	self.Connection.Write(content)
}
