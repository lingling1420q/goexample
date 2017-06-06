package main

import (
	"log"
	"strconv"
)

func main() {
	//string to int
	num, _ := strconv.Atoi("8888")
	log.Println(num)

	//int to string
	log.Println(strconv.Itoa(9000))

}
