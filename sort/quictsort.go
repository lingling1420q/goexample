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

func main() {
	a := []int{3, 1, 5, 7, 2, 4, 9, 6, 10, 8}
	fmt.Println(a)
	quickSort(a, 0, 9)
	fmt.Println(a)
}
