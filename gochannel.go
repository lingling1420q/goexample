package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	fmt.Println("1")
	channel <- "h1"
	fmt.Println("2")
	channel <- "w2"
	// fmt.Println("3")
	// channel <- "c3" // 执行到这一步，直接报 error
	// fmt.Println("...")
	// msg1 := <-channel
	// fmt.Println(msg1)

	for item := range channel {
		fmt.Println("item:", item)
	}
	close(channel)
}
