package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	c := w.(http.Response)
	body := req.FormValue("name")
	fmt.Println(body)
	io.WriteString(w, body)
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
