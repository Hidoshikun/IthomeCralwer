package engine

import (
	"ithome/cralwer/model"
	"log"
)

// ConcurrentEngine here
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan model.DeviceList
}

// Scheduler here
type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	ReadyNotifier
	Run()
}

// ReadyNotifier here
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

// Run here
func (e *ConcurrentEngine) Run(Seeds []Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range Seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			item1 := item.(model.DeviceList)
			log.Printf("Engine: got item from: %v, devices num: %v", item1.ArticleID, len(item1.DeviceList))
			go func(i model.DeviceList) { e.ItemChan <- i }(item1)
		}

		for _, request := range result.Requests {
			if isDuplicate(request.URL) != true {
				e.Scheduler.Submit(request)
			}
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false
}
