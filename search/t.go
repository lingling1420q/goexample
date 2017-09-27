package main

import "fmt"

func sort(arr []int, n int) {
	tmp := make([]int, len(arr)+1)
	for _, v := range arr {
		if v > len(tmp) {
			l := v - len(tmp)
			tmp = append(tmp, make([]int, l)...)
		}
		tmp[v] = 1
	}
	for index, v := range tmp {
		if v == 1 && n-index < len(tmp) && tmp[n-index] == 1 {
			fmt.Println(index, "+", n-index, "=", n)
		}
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	sort(arr, 20)
}
