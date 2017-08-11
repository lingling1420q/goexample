package main

import (
	"fmt"
	"net/url"
)

func main() {
	var s string

	s = "****"
	s = "r/wW3wsbr_QV=rBeDc4in"
	s = url.QueryEscape(s)

	fmt.Println(s)

	a, _ := url.QueryUnescape(s)
	fmt.Println(a)
}
