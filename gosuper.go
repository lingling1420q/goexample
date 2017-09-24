package main

import (
	"fmt"
)

type Operation interface {
	Say(content string)
}

type A struct {
	Text string
}

type B struct {
	A
}

func (self *A) Say(content string) {
	fmt.Println("a say:", self.Text)
}

// func (self *B) Say(content string) {
// 	self.A.Say("content")
// 	fmt.Println("b say:", content)
// }

func main() {
	var b Operation
	b = &B{}

	b.A.Text = "ceshi"
	b.Say("content")
}

// package main
// import (
//     "fmt"
// )
// type A struct {
//     Text string
// }
// type Operator interface {
//     Say()
// }
// func (a *A) Say() {
//     fmt.Printf("A::Say():%s\n", a.Text)
// }
// type B struct {
//     A
// }
// func (b *B) Say() {
//     b.A.Say()
//     fmt.Printf("B::Say():%s\n", b.Text)
// }
// func main() {
//     b := B{}
//     b.Text = "hello,world"
//     b.Say()
// }
