package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	//'{"username": "dp_peng001", "password": "r/wW3wsbr_QV=rBeDc4in"}'
	query := "password=r/wW3wsbr_QV=rBeDc4in&username=dp_peng001"
	v, err := url.ParseQuery(query)
	log.Println(err)
	log.Println(v.Encode())
	resp, err := http.PostForm("https://api3.verycloud.cn/API/OAuth/authorize", v)
	//resp, err := client.Post(u, "application/x-www-form-urlencoded", strings.NewReader(string(reqr)))
	log.Println(err)
	result, err := ioutil.ReadAll(resp.Body)
	log.Println(string(result))
}
