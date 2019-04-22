package parser

import (
	"fmt"
	"ithome/cralwer/engine"
	"ithome/cralwer/model"
	"regexp"
)

const titleRe = `<h1>(.*?)</h1>`
const timeRe = `<span id="pubtime_baidu">(.*?)</span>`

const articleListRe = `https://www.ithome.com/0/[\d]{3}/[\d]{3}.htm`

// ParseArticle here
func ParseArticle(contents []byte, _url string) engine.ParseResult {
	article := model.Article{}
	re1 := regexp.MustCompile(titleRe)
	match1 := re1.FindSubmatch(contents)
	if match1 != nil {
		title := string(match1[1])
		fmt.Println("article title: ", title)
		article.Title = title
	}
	re2 := regexp.MustCompile(timeRe)
	match2 := re2.FindSubmatch(contents)
	if match2 != nil {
		time := string(match2[1])
		fmt.Println("article time: ", time)
		article.Time = time
	}
	article.Content = "Hello world to elastic search"

	re3 := regexp.MustCompile(articleListRe)
	matches3 := re3.FindAllString(string(contents), -1)
	result := engine.ParseResult{}
	for _, m := range matches3 {
		result.Requests = append(result.Requests,
			engine.Request{
				URL:        string(m),
				ParserFunc: ParseArticle,
			})
	}

	//result.Items = append(result.Items, article)

	return result
}
