package persist

import (
	"container/list"
	"database/sql"
	"ithome/cralwer/model"
	"log"
	"strings"

	// used import
	_ "github.com/go-sql-driver/mysql"
)

// ArticleURLSaver here
// 保存文章URL
func ArticleURLSaver() (chan model.ArticleURL, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/nlp?charset=utf8")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	out := make(chan model.ArticleURL)
	go func() {
		stmtIns, err := db.Prepare("INSERT INTO ithome_urls_2016 (time, url, has_fetched) VALUES( ?, ? , 0)")
		if err != nil {
			log.Printf("Item Saver: error Prepare %v", err)
		}
		defer stmtIns.Close()

		itemCount := 0
		for {
			articleURL := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, articleURL)
			itemCount++
			go saveArticleURL(stmtIns, articleURL)
		}
	}()

	return out, nil
}

// saveArticleURL here
func saveArticleURL(stmtIns *sql.Stmt, articleURL model.ArticleURL) {
	_, err := stmtIns.Exec(articleURL.Time, articleURL.URL)
	if err != nil {
		log.Printf("Item Saver: error Exec item %v: %v", articleURL, err)
	}
}

// CommentDeviceSaver here
// 保存评论设备情况
func CommentDeviceSaver() (chan model.DeviceList, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/nlp?charset=utf8")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	out := make(chan model.DeviceList)
	go func() {
		stmtIns, err := db.Prepare("INSERT INTO ithome_devices_2017 (article_id, devices) VALUES( ?, ?)")
		if err != nil {
			log.Printf("Item Saver: error Prepare %v", err)
		}
		defer stmtIns.Close()

		itemCount := 0
		for {
			commentDevices := <-out
			log.Printf("Item Saver: got item num: #%d: devices num: %v", itemCount, len(commentDevices.DeviceList))
			itemCount++
			go saveCommentDevices(stmtIns, commentDevices)
		}
	}()

	return out, nil
}

// saveCommentDevices here
func saveCommentDevices(stmtIns *sql.Stmt, commentDevices model.DeviceList) {
	_, err := stmtIns.Exec(commentDevices.ArticleID, strings.Join(commentDevices.DeviceList, ","))
	if err != nil {
		log.Printf("Item Saver: error Exec item %v: %v", commentDevices, err)
	}
}

// CommentSaver here
// 保存评论
func CommentSaver() (chan model.Comment, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/nlp?charset=utf8")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	out := make(chan model.Comment)
	go func() {
		stmtIns, err := db.Prepare("INSERT INTO ithome_comment_2017 (user_id, content, user_name, location, time, device, supported, aganist, article_url) VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			log.Printf("Item Saver: error Prepare %v", err)
		}
		defer stmtIns.Close()

		itemCount := 0
		for {
			comment := <-out
			queue := list.New()
			// 入队, 压栈
			queue.PushBack(comment)
			// 出队
			i1 := queue.Front()
			queue.Remove(i1)

			log.Printf("Item Saver: got item #%d: %v", itemCount, comment.UserName)
			itemCount++
			go saveComment(stmtIns, comment)
		}
	}()

	return out, nil
}

// saveComment here
func saveComment(stmtIns *sql.Stmt, comment model.Comment) {
	_, err := stmtIns.Exec(comment.UserID, comment.Content, comment.UserName, comment.Location, comment.Time, comment.Device, comment.Supported, comment.Aganist, comment.ArticleURL)
	if err != nil {
		log.Printf("Item Saver: error Exec item %v: %v", comment, err)
	}
}
