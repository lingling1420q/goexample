package main

import (
	"fmt"
	"net/url"
)

func main() {
	var s string

	s = "****"
	s = "f=header_%40input%40btn_search"
	s = url.QueryEscape(s)

	fmt.Println(s)

	a, _ := url.QueryUnescape(s)
	fmt.Println(a)
}
