package main

import (
	"fmt"
)

func main() {
	arr := make([]int, 10)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	fmt.Println(arr)
	arr = append(arr, 1)
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
}
