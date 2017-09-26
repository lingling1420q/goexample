package main

import "fmt"

func sort(arr []int, max int) []int {
	a := make([]int, max)
	var result []int
	for _, i := range arr {
		a[i] = 1
	}
	for index, item := range a {
		if item == 1 {
			result = append(result, index)
		}
	}
	return result
}

func main() {

	arr := []int{4, 8, 9, 12, 3, 2, 1, 19, 23}

	fmt.Println(sort(arr, 24))
}
