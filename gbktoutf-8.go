package main

import (
	"fmt"
	"github.com/axgle/mahonia"
)

func main() {
	// 将UTF-8转换为 GBK
	gbk := mahonia.NewEncoder("gbk").ConvertString("hello,世界")
	fmt.Println(gbk)

	// "你好，世界！"的GBK编码
	gbkBytes := []byte{0xC4, 0xE3, 0xBA, 0xC3, 0xA3, 0xAC, 0xCA, 0xC0, 0xBD, 0xE7, 0xA3, 0xA1}

	// 将GBK转换为UTF-8
	utf8 := mahonia.NewDecoder("gbk").ConvertString(string(gbkBytes))
	fmt.Println(utf8)
}
