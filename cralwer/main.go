package main

import (
	"database/sql"
	"ithome/cralwer/engine"
	"ithome/cralwer/parser"
	"ithome/cralwer/persist"
	"ithome/cralwer/scheduler"
	_ "strconv"

	// used import
	_ "github.com/go-sql-driver/mysql"
)

/***
// Fetch URL
func main() {
	// start a ItemSaver to save items
	itemChan, err := persist.ArticleURLSaver()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 2,
		ItemChan:    itemChan,
	}
	Seeds := []engine.Request{}

	for i := 6280; i < 6460; i++ {
		Seeds = append(Seeds, engine.Request{
			URL:        "https://www.ithome.com/list/list_" + strconv.Itoa(i) + ".html",
			ParserFunc: parser.ParseArticleList,
		})
	}
	e.Run(Seeds)
}
***/

// Fetch comment devices
func main() {
	itemChan, err := persist.CommentDeviceSaver()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	Seeds := getSeed()
	e.Run(Seeds)
}

func getSeed() []engine.Request {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/nlp?charset=utf8")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT url FROM ithome_urls_2016")
	if err != nil {
		panic(err)
	}

	seeds := []engine.Request{}
	for rows.Next() {
		var url string
		rows.Columns()
		err = rows.Scan(&url)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, engine.Request{
			URL:        url,
			ParserFunc: parser.ParseCommentDevices})
	}
	return seeds
}
