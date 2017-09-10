package main

import (
	"./httputil"
	// "bytes"
	//"fmt"
	// "io"
	// "net"
	// "net/http"
	// "time"
)

func hander(req *httputil.HTTPRequest) {
	// body := "hello world !!!"
	// c := fmt.Sprintf("HTTP/1.1 200 msg\r\nContent-Length: %d\r\nServer: golang/server\r\nContent-Type: text/html; charset=UTF-8\r\nDate: Mon, 28 Aug 2017 15:51:38 CST\r\n\r\n%s", len(body), body)
	// fmt.Println(req.String())
	// req.Finish([]byte(c))
	handler := httputil.NewRequestHandler(req, nil)
	//fmt.Println(handler)
	//handler.Finish(map[string]string{"msg": "hellow world!"})
	result := handler.Request.Arguments
	handler.Finish(map[string]interface{}{"result": "成功", "data": result, "files": handler.Request.Files})
}

func main() {
	server := httputil.HttpServer{Port: 8888, Host: "0.0.0.0", Callback: hander}
	server.Listen()
}
