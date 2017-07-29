package main

import (
	"fmt"
	"math/rand"
)

func RandInt(min, max int64) int64 {
	//rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}

	for i := 0; i < 100; i++ {
		fmt.Println(RandInt(100, 200), ",")
	}
}
