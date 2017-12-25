package main

import (
	"fmt"
)

//问题，一共10个台阶，每次可以走1个或者2个台阶，问一共有多少种走法
func version01(k int) (n int) {

	if k == 1 {
		return 1
	}
	if k == 2 {
		return 2
	}

	return version01(k-1) + version01(k-2)
}

func main() {
	fmt.Println(version01(10))
}
