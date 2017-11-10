package main

import (
	"fmt"
	"reflect"
)

func print(t interface{}) {
	fmt.Println(t)
	fmt.Println(reflect.TypeOf(t))
}

func main() {
	arr := &[5]int{1, 2, 3, 4, 5}
	print(arr)
	s := arr[1:3]
	print(s)
	arr[1] = 10
	fmt.Println(arr, s)

	a := new([5]int)
	print(a)
	b := make([]int, 2, 5)
	print(b)
}
