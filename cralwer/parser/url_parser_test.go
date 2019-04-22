package parser

import (
	"fmt"
	"ithome/cralwer/fetcher"
	"ithome/cralwer/model"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestUrlParser(t *testing.T) {
	url := "https://www.ithome.com/list/list_5.html"
	html, err := fetcher.Fetch(url)
	if err != nil {
		panic(err)
	}

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(html)))
	if err != nil {
		panic(err)
	}

	articleURL := model.ArticleURL{}
	dom.Find("#wrapper > div.content.fl > div > ul > li").
		Each(func(i int, selection *goquery.Selection) {
			content, err := selection.Html()
			if err != nil {
				t.Logf("test error %v", err)
			}
			articleURL.URL = ReMatch(articleURLRe, content)
			articleURL.Time = ReMatch(articleTimeRe, content)
			fmt.Println(articleURL)
		})

}
