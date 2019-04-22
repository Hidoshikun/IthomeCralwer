package tools

import (
	"fmt"
	"testing"
)

func TestUrlParser(t *testing.T) {
	url1 := "https://www.ithome.com/html/game/340826.htm"
	url2 := "https://www.ithome.com/0/420/138.htm"
	url3 := "https://www.ithome.com/html/win10/384893.htm"

	fmt.Println(GetArticleID(url1))
	fmt.Println(GetArticleID(url2))
	fmt.Println(GetArticleID(url3))

}
