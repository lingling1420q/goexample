package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	//"strings"
)

var (
	APIURL string = "https://api3.verycloud.cn"
)

func index(w http.ResponseWriter, req *http.Request) {
	u := fmt.Sprintf("%s%s", APIURL, req.URL.Path)
	log.Println("url:", u)
	reqr, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	log.Println("reqr", string(reqr))
	value := ""
	rs := string(reqr)
	var f interface{}
	e := json.Unmarshal(reqr, &f)
	if e == nil {
		for k, v := range f.(map[string]interface{}) {
			log.Println(k, v)
			value = fmt.Sprintf("%s%s=%v&", value, k, v)
		}
	} else {
		value = rs
	}
	log.Println("post data:", value)
	vv, err := url.ParseQuery(value)
	resp, err := http.PostForm(u, vv)
	// req.URL = r.URL
	// log.Println(req.URL)
	var msg string
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		msg = "{\"code\": 1,\"msg\": \"failed\"}"
	} else {
		msg = bytes.NewBuffer(result).String()
	}
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Write([]byte(msg))
}

func main() {
	//规则1
	http.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
