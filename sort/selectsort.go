package main

import (
	"fmt"
)

func selectMinKey(n []int, index int) (min int) {
	min = index
	for j := index; j < len(n); j++ {
		if n[j] < n[min] {
			min = j
		}
	}
	return
}

func SelectSortV1(nums []int) {
	size := len(nums)
	fmt.Println(nums, size)
	for i := 0; i < size; i++ {
		key := selectMinKey(nums, i)
		//print key
		if key != i {
			nums[key], nums[i] = nums[i], nums[key]
		}
		fmt.Println(nums, key, i)
	}
}

func main() {
	a := []int{3, 1, 5, 7, 2, 4, 9, 6}
	SelectSortV1(a)
}
