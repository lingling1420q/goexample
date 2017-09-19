package main

import "fmt"

func and(a []int, b []int) []int {
	var result []int
	aSize := len(a)
	bSize := len(b)
	i, j := 0, 0
	for i < aSize && j < bSize {
		if a[i] > b[j] {
			j += 1
		} else if a[i] < b[j] {
			i += 1
		} else if a[i] == b[j] {
			result = append(result, a[i])
			j += 1
			i += 1
		}
	}
	return result
}

func and2(a []int, b []int) []int {
	var result []int
	aSize := len(a)
	bSize := len(b)
	i, j := 0, 0
	var font int
	for i < aSize && j < bSize {
		if a[i] > b[j] {
			j += 1
		} else if a[i] < b[j] {
			i += 1
		} else if a[i] == b[j] {
			if font != a[i] {
				result = append(result, a[i])
				font = a[i]
			}
			j += 1
			i += 1
		}
	}
	return result
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 6, 8, 9}
	b := []int{6, 6, 8, 9, 10, 11, 12}

	fmt.Println(and(a, b))
	fmt.Println(and2(a, b))
}
