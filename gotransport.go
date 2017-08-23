package main

import (
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

type Proxy struct {
	ProxyAdr   string
	BackendAdr string
	Listener   net.Listener
}

func (httpproxy *Proxy) Run() {
	log.Println("server listens on ", httpproxy.ProxyAdr)
	err := http.ListenAndServe(httpproxy.ProxyAdr, httpproxy)
	log.Println("err:", err)
}

var transport = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	ResponseHeaderTimeout: 30 * time.Second,
}

func (httpproxy *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "CONNECT" {
		hij, ok := w.(http.Hijacker)
		if !ok {
			log.Println("httpserver does not support hijacking")
			return
		}
		proxyclient, _, err := hij.Hijack()
		if err != nil {
			log.Println("Cannot hijack connection " + err.Error())
			return
		}
		httpproxy.serve(proxyclient, r)
		return
	}
	r.RequestURI = ""

	resp, err := transport.RoundTrip(r)
	if err != nil {
		log.Println("RoundTrip", err)
		return
	}
	defer resp.Body.Close()

	copyHeaders(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func (httpproxy *Proxy) serve(client net.Conn, r *http.Request) {

	server, err := net.Dial("tcp", r.Host)
	if err != nil {
		log.Println(err)
		return
	}

	success := []byte("HTTP/1.1 200 Connection established\r\n\r\n")
	_, err = client.Write(success)
	if err != nil {
		log.Println(err)
		return
	}

	go func() {
		io.Copy(client, server)
		server.Close()
		client.Close()
	}()
	io.Copy(server, client)
	client.Close()
	server.Close()
}

func copyHeaders(dst, src http.Header) {
	for k, vs := range src {
		for _, v := range vs {
			dst.Add(k, v)
		}
	}
}

func main() {
	var proxyAdr string
	flag.StringVar(&proxyAdr, "a", "", "Network host to listen on.")
	flag.Parse()
	var p *Proxy
	p = &Proxy{ProxyAdr: proxyAdr}
	p.Run()
}
