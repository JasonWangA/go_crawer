package scheduler

import (
	"../engine"
)

type SimpleScheduler struct {
    workerChan chan  engine.Request
}


func (s *SimpleScheduler) ConfigWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request)  {
	go func() {
		s.workerChan <- r
	}()
}


