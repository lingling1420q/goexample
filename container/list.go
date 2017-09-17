package main

import (
	"container/list"
	"fmt"
	//"container/list"
)

func print(l list.List) {
	result := []interface{}{}
	e := l.Front()
	result = append(result, e.Value)
	for {
		e = e.Next()
		if e == nil {
			break
		}
		result = append(result, e.Value)
	}
	fmt.Println("print:", result)
}

func main() {
	var l list.List

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	fmt.Println("l:", l)
	print(l)
	l.PushFront(0)
	l.PushFront("aa")
	print(l)
}
