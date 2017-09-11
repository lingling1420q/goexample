package tcp

import (
	//logs "github.com/yangaowei/gologs"
	"testing"
	"time"
)

func TestBase(t *testing.T) {
	go server()
	//time.Sleep(2 * time.Second)
	clinet()
	time.Sleep(30 * time.Second)
}
