package main

import (
	"fmt"
	"net/url"
)

func main() {
	var s string

	s = "mongodb://waqu_crawl:crawl@123457@114.55.108.168:27017"

	s = url.QueryEscape(s)

	fmt.Println(s)

	a, _ := url.QueryUnescape(s)
	fmt.Println(a)
}
