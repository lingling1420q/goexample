package rabbitmq

import (
	//"fmt"
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	SendText("yangaowei")
	go ReceiveText()
	time.Sleep(time.Second * 3)
	SendText("yangaowei")
}
