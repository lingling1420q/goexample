package main

import (
	"fmt"
)

func defer1() (r int) {
	defer func() {
		fmt.Println("def")
		r++
	}()
	defer func() {
		fmt.Println("def2")
		r++
	}()
	return 10
}

// < == >
// func f() (result int) {
//      result = 0  //return语句不是一条原子调用，return xxx其实是赋值＋ret指令
//      func() { //defer被插入到return之前执行，也就是赋返回值和ret指令之间
//          result++
//      }()
//
//      return
// }

func defer2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

// < == >
// func defer2() (r int) {
//      t := 5
//      r = t //赋值指令
//      func() {        //defer被插入到赋值与返回之间执行，这个例子中返回值r没被修改过
//          t = t + 5
//      }
//      return        //空的return指令
// }

func defer3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

// < == >
// func defer3() (r int) {
//      r = 1  //给返回值赋值
//      func(r int) {        //这里改的r是传值传进去的r，不会改变要返回的那个r值
//           r = r + 5
//      }(r)
//      return        //空的return
// }

func main() {
	fmt.Println(defer1()) //11
	fmt.Println(defer2()) //5
	fmt.Println(defer3()) //1
}

//defer是在return之前执行的。这个在 官方文档中是明确说明了的。要使用defer时不踩坑，最重要的一点就是要明白，return xxx这一条语句并不是一条原子指令!
// 返回值 = xxx
// 调用defer函数
// 空的return
