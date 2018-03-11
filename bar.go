package main

import (
	"fmt"
	"github.com/cnych/starjazz/mathx"
	"strconv"
	"time"
)

func main() {
	for i := 1; i <= 100; i += 1 {
		str := "[" + bar(i, 100) + "] " + strconv.Itoa(i) + "%"
		fmt.Printf("\r%s", str)
		time.Sleep(1 * time.Second)
	}
	fmt.Println()
}

func bar(count, size int) string {
	str := ""
	for i := 0; i < size; i++ {
		if i < count {
			str += "="
		} else {
			str += "_"
		}
	}
	return str
}
