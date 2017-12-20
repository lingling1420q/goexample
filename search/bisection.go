package main

import (
	"fmt"
)

func bisectSearch(arr []int, k int) int {
	low := 0
	height := len(arr) - 1
	for low <= height {
		mid := (low + height) / 2
		sub := arr[mid]
		if sub > k {
			height = mid - 1
		} else if sub < k {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1

}

func main() {
	arr := []int{5, 7, 12, 14, 17, 20, 30, 41, 50, 52, 55, 57, 70, 90, 100, 101}
	fmt.Println(bisectSearch(arr, 20))
}
