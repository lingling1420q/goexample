// Tests for gocron
package gocron

import (
	"log"
	"testing"
	"time"
)

var err = 1

func task() {
	log.Println("I am a running job.")
	time.Sleep(3 * time.Second)
}

func taskWithParams(a int, b string) {
	log.Println(a, b)
	time.Sleep(3 * time.Second)
}

func TestSecond(*testing.T) {
	defaultScheduler.Every(1).Second().Do(task)
	defaultScheduler.Every(1).Second().Do(taskWithParams, 1, "hello")
	<-defaultScheduler.Start()

	defaultScheduler.Every(1).Second().Do(taskWithParams, 2, "hello")
	<-defaultScheduler.Start()
}
