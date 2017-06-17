package main

import (
	"fmt"
)

func main() {
	a := []string{}
	a = append(a, "c")
	a = append(a, "d")
	a = append(a, "a")

	fmt.Println(a)
}
