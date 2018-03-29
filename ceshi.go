package main

import (
	"fmt"
	"time"
)

type query func(string) string

func exec(name string, vs ...query) string {
	ch := make(chan string, len(vs))
	fn := func(i int) {
		ch <- vs[i](name)
	}
	for i, _ := range vs {
		go fn(i)
	}
	// var s string
	// for i, _ := range vs {
	// 	fmt.Println(i)
	// 	s += <-ch
	// }
	//<-ch
	time.Sleep(1 * time.Second)
	fmt.Println(len(ch))
	return `aaa`
}

func main() {
	ret := exec("111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	}, func(n string) string {
		return n + "func5"
	})
	fmt.Println(ret)
}
