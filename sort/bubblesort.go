package main

import (
	"fmt"
)

func bubbleSortV1(num []int) {
	size := len(num)
	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			if num[j] > num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
			}
		}
		fmt.Println(num, i)
	}
}

func bubbleSortV2(num []int) {
	size := len(num)
	i := size - 1
	for i < size && i > 0 {
		pos := 0
		for j := 0; j < i; j++ {
			if num[j] > num[j+1] {
				pos = j
				num[j], num[j+1] = num[j+1], num[j]
			}
		}
		i = pos
		fmt.Println(num, i)
	}
}

func main() {
	a := []int{3, 1, 5, 7, 2, 4, 9, 6}
	bubbleSortV1(a)
	fmt.Println()
	b := []int{3, 1, 5, 7, 2, 4, 9, 6}
	bubbleSortV2(b)
}
