package main

import "fmt"

func main() {
	x := 1
	defer func(a *int) {
		fmt.Println("a =", *a)
	}(&x)
	defer func(b int) {
		fmt.Println("b =", b)
	}(x)
	defer func() {
		fmt.Println("x =", x)
	}()

	x++
}
