package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
)

func TF(a int, b string) string {
	var result string
	result = strconv.Itoa(a) + b
	return result
}

func main() {
	fmt.Println(TF(1, `2`))
	f := reflect.ValueOf(TF)
	fmt.Println(f)
	fmt.Println(f.Type().NumIn())
	in := make([]reflect.Value, f.Type().NumIn())
	params := []interface{}{1, "2"}
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result := f.Call(in)

	fmt.Println(result[0].String())

	fmt.Println(runtime.FuncForPC(reflect.ValueOf(TF).Pointer()).Name())
}
