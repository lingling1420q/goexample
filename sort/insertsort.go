package main

import (
	"fmt"
)

// 二、直接插入排序
// 将一个记录插入到已经排好序的有序表中, 从而得到一个新的,记录数增1的有序表
// 时间复杂度也为O(n^2), 比冒泡法和选择排序的性能要更好一些
//直接插入排序(Straight Insertion Sort)

func StraighInsertSort(nums []int) {
	size := len(nums)
	fmt.Println(nums, size, 0)
	for i := 1; i < size; i++ {
		if nums[i] < nums[i-1] {
			j := i
			x := nums[i]
			for j > 0 && x < nums[j-1] {
				nums[j] = nums[j-1]
				j = j - 1
			}
			nums[j] = x
		}
		fmt.Println(nums, size, i)
	}
}

func Sort(nums []int) {
	size := len(nums)

	for i := 1; i < size; i++ {
		if nums[i] < nums[i-1] {
			j := i
			x := nums[i]
			for j > 0 && x < nums[j-1] {
				nums[j] = num[j-1]
				j -= 1
			}
			nums[j] = x
		}
	}
}

//输出
// [3 1 5 7 2 4 9 6] 8 0
// [1 3 5 7 2 4 9 6] 8 1
// [1 3 5 7 2 4 9 6] 8 2
// [1 3 5 7 2 4 9 6] 8 3
// [1 2 3 5 7 4 9 6] 8 4
// [1 2 3 4 5 7 9 6] 8 5
// [1 2 3 4 5 7 9 6] 8 6
// [1 2 3 4 5 6 7 9] 8 7

func main() {
	a := []int{3, 1, 5, 7, 2, 4, 9, 6}
	StraighInsertSort(a)
}
