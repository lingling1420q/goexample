package main

import (
	"log"
	"strings"
)

func main() {
	str := "yangaowei"
	log.Println(str[:len(str)-1])

	values := strings.SplitN("/yan", "?", 2)

	log.Println(values)
}
