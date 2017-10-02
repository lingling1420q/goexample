package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func PrintLocalDial(network, addr string) (net.Conn, error) {
	dial := net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	conn, err := dial.Dial(network, addr)
	if err != nil {
		return conn, err
	}

	fmt.Println("connect done, use", conn.LocalAddr().String())

	return conn, err
}

var (
	client = &http.Client{
		Transport: &http.Transport{
			Dial: PrintLocalDial,
			//MaxIdleConnsPerHost: 10,  默认为2
		},
	}
)

func doGet(client *http.Client, url string, id int) {
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	buf, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%d: %d -- %v\n", id, len(buf), err)
	if err := resp.Body.Close(); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v: done\n", id)
}

func doGet1(client *http.Client, url string, id int) {
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	// buf, err := ioutil.ReadAll(resp.Body)
	// fmt.Printf("%d: %d -- %v\n", id, len(buf), err)
	if err := resp.Body.Close(); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v: done\n", resp)
}

func main() {
	const URL = "http://www.baid.com/"
	//const URL = "http://120.26.13.218:8888/"

	for {
		go doGet(client, URL, 1)
		go doGet(client, URL, 2)
		go doGet(client, URL, 3)
		time.Sleep(2 * time.Second)
	}
}
