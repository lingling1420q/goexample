package main

import (
	"flag"
	"log"
)

var (
	name  string
	age   int
	marry bool
)

// -flag
// -flag=x
// -flag x  // 只有非boolean标签能这么用
func init() {
	flag.StringVar(&name, "name", "yangaowei", "help name")
	flag.IntVar(&age, "age", 18, "help age")
	flag.BoolVar(&marry, "marry", false, "help marry")
	flag.Parse()
}

func main() {
	log.Printf("name:%s; age:%d; marry:%v\n", name, age, marry)
}

//name:asdfa; age:20; marry:true
//go run goflag.go -age=20 -name=asdfa  -marry true
