package persist

import (
	"IthomeCralwer/cralwer/model"
	"database/sql"
	"log"
)

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
