package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

//有两种信号不能被拦截和处理: SIGKILL和SIGSTOP。
func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGKILL)
	fmt.Println("...")
	// Block until a signal is received.
	s := <-c
	fmt.Println("......")
	fmt.Println("Got signal:", s)
}
