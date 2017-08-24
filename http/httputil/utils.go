package httputil

import (
	"fmt"
	logs "github.com/yangaowei/gologs"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Request struct {
	http.Request
	PostData string
}

var (
	CRLF string = "\r\n"
)

func parseRequestStartLine(startLine string) (method, path, version string) {
	logs.Log.Debug("startLine %#s", startLine)
	V := strings.Split(startLine, " ")
	method, path, version = V[0], V[1], V[2]
	return
}

func parseHeader(headerLine string) (header http.Header) {
	header = make(http.Header)
	logs.Log.Debug("headerLine %#s", headerLine)
	headerLine = strings.TrimLeft(headerLine, `\r\n`)
	for _, line := range strings.Split(headerLine, `\r\n`) {
		l := strings.SplitN(line, ":", 2)
		header.Add(l[0], strings.TrimLeft(l[1], " "))
	}
	return
}

func parseBody(bodyLine string) (body string) {
	logs.Log.Debug("bodyLine %#s", bodyLine)

	return strings.TrimLeft(bodyLine, `\r\n\r\n`)
}

func ParseRequest(content string) (req Request) {
	logs.Log.Debug("content %#s", content)
	content = strings.TrimLeft(content, "\"")
	content = strings.TrimRight(content, "\"")
	logs.Log.Debug("content %#s", content)
	if len(content) > 10 {
		startLinePoint := strings.Index(content, `\r\n`)
		endLinePoint := strings.Index(content, `\r\n\r\n`)
		method, path, version := parseRequestStartLine(content[:startLinePoint])
		header := parseHeader(content[startLinePoint:endLinePoint])
		req.Method = strings.ToUpper(method)
		req.Proto = version
		req.RequestURI = path
		req.Header = header
		logs.Log.Debug("req.Method %#s", req.Method)
		if req.Method == "POST" {
			req.PostData = parseBody(content[endLinePoint:])
			logs.Log.Debug("PostData %#s", req.PostData)
		}
	}
	return
}

func WriteHeader(startLine []string, header http.Header, chunck string) (headerString string) {
	line := []string{}
	fmt.Println("len(chunck)", len(chunck))
	header.Add("Content-Length", strconv.Itoa(len(chunck)))
	header.Add("Server", "golang/server")
	if _, ok := header["Content-Typ"]; !ok {
		header.Add("Content-Type", "text/html; charset=UTF-8")
	}
	header.Add("Date", time.Now().Format(time.RFC1123))
	line = append(line, fmt.Sprintf("HTTP/1.1 %s %s", startLine[1], startLine[2]))
	for key, value := range header {
		line = append(line, fmt.Sprintf("%s: %s", key, value[0]))
	}
	headerString = strings.Join(line, CRLF) + CRLF + CRLF
	return
}
