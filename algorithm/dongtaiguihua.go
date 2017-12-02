package main

import "fmt"

func version01(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n >= 3 {
		return version01(n-1) + version01(n-2)
	} else {
		return 0
	}
}

var (
	m = make(map[int]int)
)

func version02(n int) int {
	if val, ok := m[n]; ok {
		return val
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n >= 3 {
		value := version02(n-1) + version02(n-2)
		m[n] = value
		return value
	} else {
		return 0
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func version03(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	a, b := 1, 2
	tmp := 0
	for i := 3; i <= n; i++ {
		tmp = a + b
		a = b
		b = tmp
	}
	return tmp
}

var (
	kuang = map[int]int{500: 5, 200: 3, 300: 4, 350: 3, 400: 5}
	p     = []int{5, 3, 4, 3, 5}
	g     = []int{500, 200, 300, 350, 400}
)

func kuang01(n, w int) int {
	// keys := []int{}
	// for key, _ := range kuang {
	// 	keys = append(keys, key)
	// }
	// fmt.Println(keys)
	if n <= 1 && w < p[0] {
		return 0
	}
	if n == 1 && w >= p[0] {
		return g[0]
	}
	if n > 1 && w < p[n-1] {
		return kuang01(n-1, w)
	}

	return max(kuang01(n-1, w), kuang01(n-1, w-p[n-1])+g[n-1])
}

func main() {

	fmt.Println(version01(10))
	fmt.Println(version02(10))
	fmt.Println(version03(10))
	fmt.Println(max(1, 2))

	fmt.Println(kuang01(5, 10))
}
