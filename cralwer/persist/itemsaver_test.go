package persist

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestSQL(t *testing.T) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/nlp?charset=utf8")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT url FROM ithome_urls")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var url string
		rows.Columns()
		err = rows.Scan(&url)
		if err != nil {
			panic(err)
		}
		fmt.Println(url)
	}
}
