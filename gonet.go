package main

import (
	"fmt"
	"net"
)

func main() {
	host, _ := net.LookupHost("www.baidu.com")
	fmt.Println(host)
}
