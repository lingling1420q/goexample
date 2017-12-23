package main

import (
	"fmt"
)

func sumK(arr []int, k int) {
	low := 0
	high := len(arr) - 1
	for low < high {
		h := arr[low] + arr[high]
		if h > k {
			high -= 1
		} else if h < k {
			low += 1
		} else {
			fmt.Println(arr[low], arr[high])
			low += 1
			high -= 1
		}
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	sumK(a, 20)
}
