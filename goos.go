package main

import (
	"log"
	"os"
)

func main() {
	f, _ := os.Open("gojson.go")
	b1 := make([]byte, 10)
	for {
		n1, err := f.Read(b1)
		log.Println(n1, string(b1))
		if err != nil {
			break
		}
	}
}
