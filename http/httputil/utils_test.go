package httputil

import (
	"fmt"
	logs "github.com/yangaowei/gologs"
	"net/http"
	"testing"
	//"time"
)

func TestBase(t *testing.T) {
	content := `POST /API/dns/planList HTTP/1.1\r\nHost: 120.26.13.218:801\r\nConnection: keep-alive\r\nContent-Length: 38\r\nPostman-Token: 7254ba5d-952c-4f73-29c1-680bd554ba07\r\nCache-Control: no-cache\r\nOrigin: chrome-extension://fhbjgbiflinjbdggehcddcbncdddomop\r\nUser-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36\r\nContent-Type: application/x-www-form-urlencoded\r\nAccept: */*\r\nAccept-Encoding: gzip, deflate\r\nAccept-Language: en-US,en;q=0.8,zh-CN;q=0.6,zh;q=0.4,vi;q=0.2,zh-TW;q=0.2\r\n\r\ntoken=db2ed3f770658ff5fcb0050826f8db37`
	req := ParseRequest(content)
	for key, value := range req.Header {
		logs.Log.Debug("%s:%s", key, value)
	}

	// "Server": "TornadoServer/%s" % tornado.version,
	//         "Content-Type": "text/html; charset=UTF-8",
	//         "Date": httputil.format_timestamp(time.time()),
	fmt.Println("-------------------------------")
	header := make(http.Header)
	header.Add("Server", "golang")
	header.Add("Content-Type", "text/html; charset=UTF-8")
	header.Add("Date", "Sun, 27 Jan 2013 18:43:20 GMT")
	headerStr := WriteHeader([]string{"", "200", "msg"}, header, "helloword")

	fmt.Println("headerStr:", headerStr)
}

func hander(req *HTTPRequest) {
	// body := "hello world !!!"
	// c := fmt.Sprintf("HTTP/1.1 200 msg\r\nContent-Length: %d\r\nServer: golang/server\r\nContent-Type: text/html; charset=UTF-8\r\nDate: Mon, 28 Aug 2017 15:51:38 CST\r\n\r\n%s", len(body), body)
	// fmt.Println(req.String())
	// req.Finish([]byte(c))
	handler := NewRequestHandler(req, nil)
	//handler.Finish(map[string]string{"msg": "hellow world!"})
	result := handler.request.Arguments
	handler.Finish(map[string]interface{}{"result": "success", "data": result})
}

func TestHttpServer(t *testing.T) {
	server := HttpServer{Port: 8888, Host: "0.0.0.0", Callback: hander}
	server.Listen()
}
