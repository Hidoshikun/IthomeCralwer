package main

import (
	"IthomeCralwer/cralwer/config"
	"IthomeCralwer/cralwer/engine"
	"IthomeCralwer/cralwer/parser"
	"IthomeCralwer/cralwer/persist"
	"IthomeCralwer/cralwer/scheduler"
	"IthomeCralwer/cralwer/tools"

	"database/sql"
	"strconv"

	// used import
	_ "github.com/go-sql-driver/mysql"
)

// Fetch comment devices
func main() {
	fetchComment()
}

// fetchComment
// 获取所有评论接口
func fetchComment() {
	itemChan, err := persist.CommentDeviceSaver()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 200,
		ItemChan:    itemChan,
	}
	Seeds := getURLSeed()
	e.Run(Seeds)
}

// fetchArticleURL
// 获取所有文章URL
func fetchArticleURL() {
	// start a ItemSaver to save items
	itemChan, err := persist.CommentDeviceSaver()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 2,
		ItemChan:    itemChan,
	}
	Seeds := []engine.Request{}

	for i := 250; i < 1300; i++ {
		Seeds = append(Seeds, engine.Request{
			URL:        "https://www.ithome.com/list/list_" + strconv.Itoa(i) + ".html",
			ParserFunc: parser.ParseArticleList,
		})
	}
	e.Run(Seeds)
}

// getURLSeed
// 从数据库取符合条件的URL，生成engine的种子传入
func getURLSeed() []engine.Request {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/nlp?charset=utf8")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT url FROM ithome_urls_" + config.Year)
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

		articleID := tools.GetArticleID(urla)
		urlb := "https://m.ithome.com/api/comment/newscommentlistget?NewsID=" + articleID + "&LapinID=&MaxCommentID=0&Latest="
		seeds = append(seeds, engine.Request{
			URL:        urlb,
			ParserFunc: parser.ParseCommentDevices})
	}
	return seeds
}
