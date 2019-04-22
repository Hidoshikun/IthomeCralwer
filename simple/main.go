package main

import (
	"ithome/cralwer/engine"
	"ithome/cralwer/parser"
	"ithome/cralwer/persist"
	"ithome/cralwer/scheduler"
	"strconv"
)

// Fetch URL
func main() {
	// start a ItemSaver to save items
	itemChan, err := persist.ArticleURLSaver()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan,
	}
	Seeds := []engine.Request{}

	for i := 1; i < 100; i++ {
		Seeds = append(Seeds, engine.Request{
			URL:        "https://www.ithome.com/list/list_" + strconv.Itoa(i) + ".html",
			ParserFunc: parser.ParseArticleList,
		})
	}
	e.Run(Seeds)
}
