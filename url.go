package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	var s string

	s = "****"
	s = "f=header_%40input%40btn_search"
	s = url.QueryEscape(s)

	fmt.Println(s)

	a, _ := url.QueryUnescape("\"")
	fmt.Println(a)

	m, err := url.ParseQuery(`/url?x=1&y=2&y=3;z`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
}
