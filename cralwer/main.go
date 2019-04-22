package main

import (
	"database/sql"
	"ithome/cralwer/engine"
	"ithome/cralwer/parser"
	"ithome/cralwer/persist"
	"ithome/cralwer/scheduler"
	"ithome/cralwer/tools"
	_ "strconv"

	// used import
	_ "github.com/go-sql-driver/mysql"
)

const articleIDRe = `/([\d]*).htm`

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
		WorkerCount: 4,
		ItemChan:    itemChan,
	}
	Seeds := []engine.Request{}

	for i := 2200; i < 3350; i++ {
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
		WorkerCount: 200,
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

	rows, err := db.Query("SELECT url FROM ithome_urls_2017")
	if err != nil {
		panic(err)
	}

	seeds := []engine.Request{}
	for rows.Next() {
		var urla string
		rows.Columns()
		err = rows.Scan(&urla)
		if err != nil {
			panic(err)
		}
		articleID := tools.ReSubMatch(articleIDRe, urla)
		urlb := "https://m.ithome.com/api/comment/newscommentlistget?NewsID=" + articleID + "&LapinID=&MaxCommentID=0&Latest="
		seeds = append(seeds, engine.Request{
			URL:        urlb,
			ParserFunc: parser.ParseCommentDevices})
	}
	return seeds
}
