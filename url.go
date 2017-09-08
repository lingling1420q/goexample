package main

import (
	"fmt"
	"log"
	"net/url"
	"unicode"
)

func main() {
	var s string

	s = "****"
	s = "&#29305;&#27922;&#33073;"
	n, _ := url.Parse(s)

	fmt.Println(n)
	fmt.Println(unicode(s))
	a, _ := url.QueryUnescape("\"")
	fmt.Println(a)

	m, err := url.ParseQuery(`/url?x=1&y=2&y=3;z`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
}
