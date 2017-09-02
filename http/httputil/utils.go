package httputil

import (
	"bytes"
	"fmt"
	logs "github.com/yangaowei/gologs"
	"net/http"
	"net/url"
	"os"
	"regexp"
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

func RxOf(pattern string, content string, index int) (rcontent string) {
	re, _ := regexp.Compile(pattern)
	submatch := re.FindStringSubmatch(content)
	for i, v := range submatch {
		if i == index {
			rcontent = v
			break
		}
	}
	return
}

func parseRequestStartLine(startLine string) (method, path, version string) {
	logs.Log.Debug("startLine %#s", startLine)
	V := strings.Split(startLine, " ")
	method, path, version = V[0], V[1], V[2]
	return
}

func parseHeader(headerByte []byte) (header http.Header) {
	header = make(http.Header)
	//headerByte = headerByte[startLinePoint+len([]byte(CRLF)) : endLinePoint]
	for _, line := range bytes.Split(headerByte, []byte(CRLF)) {
		l := bytes.SplitN(line, []byte(":"), 2)
		header.Add(string(l[0]), strings.TrimLeft(string(l[1]), " "))
	}
	return
}

func encode(c string) string {
	c = strings.TrimSpace(c)
	c = strings.TrimRight(c, "\"")
	c = strings.TrimLeft(c, "\"")
	return c
}

func parseFrom(fromData []byte) map[string]string {
	result := make(map[string]string)
	for _, item := range bytes.Split(fromData, []byte(";")) {
		re, _ := regexp.Compile(`(:|=)`)
		point := bytes.Index(item, []byte(CRLF))
		if point > 0 {
			item = item[:point]
		}
		kv := re.Split(string(item), 2)
		result[encode(kv[0])] = encode(kv[1])
	}
	return result
}

func parseBody(contentType string, body []byte) (arguments url.Values, files map[string]interface{}) {
	logs.Log.Debug("body %#v", string(body))
	arguments = make(map[string][]string)
	files = make(map[string]interface{})
	if strings.Index(contentType, "multipart/form-data") == 0 {
		boundary := RxOf(`boundary=(.+)`, contentType, 1)
		endBoundaryPoint := bytes.Index(body, []byte("--"+boundary+"--"))
		for _, item := range bytes.Split(body[:endBoundaryPoint], []byte("--"+boundary+"\r\n")) {
			if len(item) == 0 {
				continue
			}
			eoh := bytes.Index(item, []byte(CRLF+CRLF))
			h := parseFrom(item[:eoh])
			name := h["name"]
			if filename, ok := h["filename"]; ok {
				temporary := fmt.Sprintf("%d-%s", time.Now().Unix(), filename)
				f, _ := os.Create(temporary)
				f.Write([]byte(encode(string(item[eoh+4 : len(item)-2]))))
				f.Close()
				files[filename] = temporary
			} else {
				arguments[name] = append(arguments[name], encode(string(item[eoh+4:len(item)-2])))
			}
		}
	}
	return
}

func parseQuery(query string) (value url.Values) {
	values := strings.SplitN(query, "?", 2)
	if len(values) == 2 {
		value, _ = url.ParseQuery(values[1])
	}

	return
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
		header := parseHeader([]byte(content[startLinePoint:endLinePoint]))
		req.Method = strings.ToUpper(method)
		req.Proto = version
		req.RequestURI = path
		req.Header = header
		logs.Log.Debug("req.Method %#s", req.Method)
		// if req.Method == "POST" {
		// 	req.PostData = parseBody(content[endLinePoint:])
		// 	logs.Log.Debug("PostData %#s", req.PostData)
		// }
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
