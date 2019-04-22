package persist

import (
	"IthomeCralwer/cralwer/model"
	"database/sql"
	"log"

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
