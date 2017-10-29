package gocron

import (
	"sort"
	"time"
)

type Scheduler struct {
	// Array store jobs
	jobs []*Job

	// Size of jobs which jobs holding.
	size int
}

func (s *Scheduler) Len() int {
	return len(s.jobs)
}

func (s *Scheduler) Swap(i, j int) {
	s.jobs[i], s.jobs[j] = s.jobs[j], s.jobs[i]
}

func (s *Scheduler) Less(i, j int) bool {
	return s.jobs[j].nextRun.After(s.jobs[i].nextRun)
}

func (s *Scheduler) getRunnableJobs() (runningJobs []*Job, n int) {
	sort.Sort(s)
	for _, job := range s.jobs {
		if job.shouldRun() {
			runningJobs = append(runningJobs, job)
			n++
		} else {
			break
		}
	}
	return
}

func (s *Scheduler) NextRun() (job *Job, t time.Time) {
	if s.Len() < 0 {
		return nil, time.Now()
	}
	sort.Sort(s)
	return s.jobs[0], s.jobs[0].nextRun
}

func (s *Scheduler) Every(interval uint64) *Job {
	job := NewJob(interval)
	s.jobs = append(s.jobs, job)
	return job
}

func (s *Scheduler) RunPending() {
	runnableJobs, n := s.getRunnableJobs()

	if n != 0 {
		for i := 0; i < n; i++ {
			runnableJobs[i].run()
		}
	}
}

func (s *Scheduler) Start() chan bool {
	stopped := make(chan bool, 1)
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				s.RunPending()
			case <-stopped:
				return
			}
		}
	}()

	return stopped
}

func NewScheduler() *Scheduler {
	return &Scheduler{[]*Job{}, 0}
}

var defaultScheduler = NewScheduler()
var jobs = defaultScheduler.jobs
