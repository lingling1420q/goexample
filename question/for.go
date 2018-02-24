package main

import (
	"fmt"
)

type Stduent struct {
	name string
	age  int
}

func main() {
	st := make(map[string]*Stduent)
	stus := []Stduent{{"yan", 12}, {"gao", 13}, {"wei", 14}}
	for _, stu := range stus {
		st[stu.name] = &stu
	}

	for k, v := range st {
		fmt.Println(k, v.name)
	}

}
