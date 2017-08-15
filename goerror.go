package main

import (
	"errors"
	"fmt"
)

func funcA() error {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("panic recover! p: %v\n", p)
		}
	}()
	return funcB()
}

func funcB() error {
	// simulation
	panic("foo")
	return errors.New("success")
}

func test() {
	err := funcA()
	if err == nil {
		fmt.Printf("err is nil\n")
	} else {
		fmt.Printf("err is %v\n", err)
	}
}

func main() {
	test()
}
