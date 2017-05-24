package main

import (
	"log"
)

type Person struct {
	Name string
	Age  int
}

type Student struct {
	Person
	Level int
}

func (p Person) Fine() (name string) {
	return p.Name
}

func main() {
	s := Student{}
	log.Println(s.Fine())
}
