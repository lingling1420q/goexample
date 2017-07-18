package main

import (
	"log"
)

func main() {
	m := map[string]string{"a": "a"}
	v, ok := m["b"]
	log.Println(v, ok)

	b := []int{1, 2, 3}

	c := b[1:]

	log.Println(c)
	if v {
		log.Println("afdsfaf")
	}
}
