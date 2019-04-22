package scheduler

import "ithome/cralwer/engine"

// SimpleScheduler here
type SimpleScheduler struct {
	workerChan chan engine.Request
}

// Submit here
func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.workerChan <- r }()
}

// Run here
func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

// WorkerChan here
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

// WorkerReady here
func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}
