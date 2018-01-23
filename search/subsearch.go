package main

import (
	"fmt"
)

func search(arr []int, v int) (index int) {
	size := len(arr)
	left := 0
	right := size - 1
	for left < right {
		middle := (left + right) / 2
		if arr[middle] == v {
			index = middle
			return
		} else if arr[middle] > v {
			right = middle - 1
		} else if arr[middle] < v {
			left = middle + 1
		}
	}
	return
}

func main() {
	arr := []int{10, 22, 33, 44, 54, 64, 74, 84, 94, 104}
	fmt.Println(search(arr, 5))
}
