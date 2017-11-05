package main

import (
	"log"
	//"math/rand"
	"time"
)

var (
	MaxWorker = 1000
	MaxQueue  = 1000
)

type Playload struct {
}

func (p *Playload) UploadS3() (e error) {
	//一通具体的业务逻辑
	// second := rand.Intn(5) * time.Second
	log.Println("second: ", 1)
	time.Sleep(1 * time.Second)
	return
}

type Job struct {
	Playload Playload
}

var JobQueue chan Job

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				job.Playload.UploadS3()
			case <-w.quit:
				return
			}
		}
	}()
}

func playloadHandler() {
	playload := Playload{}
	work := Job{Playload: playload}
	JobQueue <- work
}

type Dispatcher struct {
	WorkerPool chan chan Job
	maxWorks   int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, maxWorks: maxWorkers}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.maxWorks; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				jobChannel := <-d.WorkerPool
				jobChannel <- job
			}(job)
		}
	}
}

func main() {
	log.Println(MaxQueue, MaxWorker)
	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()
	JobQueue = make(chan Job, MaxQueue)
	for i := 0; i < 10000; i++ {
		//log.Println("handler ", i)
		playloadHandler()
	}
}
