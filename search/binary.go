package main

import (
	"fmt"
)

func binary(a int, arr []int) int {
	fmt.Println(a, arr)
	size := len(arr)
	if size == 1 {
		if arr[0] == a {
			return a
		} else {
			return 0
		}
	} else {
		sub := size / 2
		if arr[sub] > a {
			return binary(a, arr[:sub])
		} else if arr[sub] < a {
			return binary(a, arr[sub:])
		} else {
			return arr[sub]
		}
	}
	return 0
}

func main() {
	result := binary(22, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13})

	fmt.Println("result:", result)
}
