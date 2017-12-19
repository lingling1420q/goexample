package main

import "fmt"

const (
	UNSET int = iota - 1
	OFFLINE
	SERVER
	CLIENT
)

func main() {
	fmt.Println(UNSET)   // -1
	fmt.Println(OFFLINE) //0
	fmt.Println(SERVER)  //1
	fmt.Println(CLIENT)  //1
}
