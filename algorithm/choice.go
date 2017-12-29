package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	str = []string{"A", "B", "C", "D"}
)

func sum(arr []int) int {
	result := 0
	for _, num := range arr {
		result += num
	}
	return result
}

func weightChoice(weight []int) (s string) {
	rand.Seed(time.Now().UnixNano())
	total := sum(weight)
	t := rand.Intn(total)
	for index, num := range weight {
		t -= num
		if t < 0 {
			s = str[index]
			break
		}
	}
	return
}

func main() {
	weight := []int{5, 2, 3, 1}
	r := make(map[string]int)
	for i := 0; i < 100; i++ {
		a := weightChoice(weight)
		if _, ok := r[a]; ok {
			r[a] += 1
		} else {
			r[a] = 1
		}
	}
	fmt.Println(r)
}
