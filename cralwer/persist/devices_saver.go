package persist

import (
	"IthomeCralwer/cralwer/config"
	"IthomeCralwer/cralwer/model"

	"database/sql"
	"log"
	"strings"
)

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
		stmtIns, err := db.Prepare("INSERT INTO ithome_devices_" + config.Year + " (article_id, devices) VALUES( ?, ?)")
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
