package main

import (
	logs "github.com/yangaowei/gologs"
	"math/rand"
	"time"
)

var (
	MaxSize int = 100
)

type Task struct {
	name string
}

func genName() string {
	rand.Seed(time.Now().UnixNano())
	ret := ""
	for len(ret) < 8 {
		tmp := rune(97 + rand.Intn(25))
		ret += string(tmp)
	}
	return ret
}

func craeteTask() (t *Task) {
	name := genName()
	t = &Task{name: name}
	return
}

func proudce(ch chan *Task) {
	for {
		size := len(ch)
		if size == MaxSize {
			logs.Log.Informational("task is maxSize,sleep 3s")
			time.Sleep(3 * time.Second)
		} else {
			ch <- craeteTask()
			time.Sleep(1 * time.Second)
		}
	}
}

func consumer(ch chan *Task) {
	for {
		select {
		case task := <-ch:
			logs.Log.Informational("handler task name %s", task.name)
			//time.Sleep( * time.Second)
		default:
			logs.Log.Warning("no task")
			time.Sleep(10 * time.Second)
		}
	}
}

func main() {
	logs.Log.Informational(genName())
	ch := make(chan *Task, MaxSize)
	go proudce(ch)
	go consumer(ch)
	time.Sleep(1000 * time.Second)
}
