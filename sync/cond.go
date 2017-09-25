package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func test(x int) {
	cond.L.Lock() // 获取锁
	fmt.Println("wait before...")
	cond.Wait() // 等待通知  暂时阻塞
	fmt.Println(x)
	fmt.Println("wait end...")
	cond.L.Unlock() // 释放锁，不释放的话将只会有一次输出
}
func main() {
	for i := 0; i < 40; i++ {
		go test(i)
	}
	fmt.Println("start all")
	os.Stdin.Read(make([]byte, 1))
	cond.Signal() //  下发广播给所有等待的goroutine
	os.Stdin.Read(make([]byte, 1))
	cond.Broadcast()
	time.Sleep(time.Second * 60)
}
