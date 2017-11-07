package main

import (
	"fmt"
	"reflect"
)

func print(t interface{}) {
	fmt.Println(t)
	fmt.Println(reflect.TypeOf(t), reflect.ValueOf(t))
}

type T int

func main() {
	var v *T
	var i interface{}
	print(i)
	print(v)
	fmt.Println(i == nil)
	fmt.Println(v == nil)
	fmt.Println(i == v)
	i = v
	fmt.Println(i == v)
	var val interface{}
	if val == nil {
		fmt.Println("val is nil")
	} else {
		fmt.Println("val is not nil")
	}

}
