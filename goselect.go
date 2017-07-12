package main

import (
	"fmt"
	"time"
)

func main() {

	c := time.Tick(1 * time.Second)
	for {
		select {
		case m := <-c:
			fmt.Println("current:", m)
		default:
			fmt.Println("default:", time.Now())
			time.Sleep(2 * time.Second)
		}
	}
}
