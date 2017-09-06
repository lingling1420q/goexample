package main

import (
	"fmt"
)

func mergeSort(arr []int) []int {
	//fmt.Println(arr)
	size := len(arr)
	if size <= 1 {
		return arr
	}
	num := size / 2
	left := mergeSort(arr[:num])
	right := mergeSort(arr[num:])
	return merge(left, right)
}

func merge(left, right []int) (result []int) {
	var l, r int
	lSize := len(left)
	rSize := len(right)
	for l < lSize && r < rSize {
		if left[l] < right[r] {
			result = append(result, left[l])
			l += 1
		} else {
			result = append(result, right[r])
			r += 1
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

func main() {
	arr := []int{3, 1, 5, 7, 2, 4, 9, 6, 10, 8, 20}
	fmt.Println(mergeSort(arr))
}
