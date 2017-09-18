package main

import (
	"fmt"
)

func partition(a []int, low int, hight int) int {
	//fmt.Println(a, low, hight)
	privotKey := a[low]
	for low < hight {
		fmt.Println(a, low, hight)
		for low < hight && a[hight] >= privotKey {
			hight -= 1
		}
		a[low], a[hight] = a[hight], a[low]
		for low < hight && a[low] <= privotKey {
			low += 1
		}
		a[low], a[hight] = a[hight], a[low]
	}
	return low
}

func quickSort(a []int, low int, hight int) {
	if low < hight {
		privotLoc := partition(a, low, hight)
		quickSort(a, low, privotLoc-1) //递归对低子表递归排序
		quickSort(a, privotLoc+1, hight)
	}
}

func quick(arr []int) {
	x := arr[0]
	i := 0
	j := len(arr) - 1

	for i < j {
		for i < j && arr[j] >= x {
			j -= 1
		}
		arr[i] = arr[j]

		for i < j && arr[i] <= x {
			i += 1
		}
		arr[j] = arr[i]
	}
	arr[i] = x
	fmt.Println(arr)
	fmt.Println(i)
}

func main() {
	a := []int{3, 1, 5, 7, 2, 4, 9, 6, 10, 8}
	quick(a)
	// fmt.Println(a)
	// quickSort(a, 0, 9)
	// fmt.Println(a)
}
