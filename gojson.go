package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func origin() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println("s:", s)
	fmt.Println("s type:", reflect.TypeOf(s))
}

//上面那种解析方式是在我们知晓被解析的JSON数据的结构的前提下采取的方案，如果我们不知道被解析的数据的格式，又应该如何来解析呢？

// 我们知道interface{}可以用来存储任意数据类型的对象，这种数据结构正好用于存储解析的未知结构的json数据的结果。JSON包中采用map[string]interface{}和[]interface{}结构来存储任意的JSON对象和数组。Go类型和JSON类型的对应关系如下：

// bool 代表 JSON booleans,
// float64 代表 JSON numbers,
// string 代表 JSON strings,
// nil 代表 JSON null.

func now() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	json.Unmarshal(b, &f)
	fmt.Println(f)
	fmt.Println(reflect.TypeOf(f))
	//那么如何来访问这些数据呢？通过断言的方式：
	m := f.(map[string]interface{})
	//通过断言之后，你就可以通过如下方式来访问里面的数据了
	fmt.Println("")
	fmt.Println("字段信息")
	for k, v := range m {
		fmt.Println(k, ":", v, ";   type:", reflect.TypeOf(v))
	}
	fmt.Println("")
	//另一种打印方式
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", reflect.TypeOf(vv))
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}

func main() {
	origin()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	now()
	//格式比较乱，可以使用fmt.printf打印
}
