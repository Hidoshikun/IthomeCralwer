package parser

import (
	"IthomeCralwer/cralwer/engine"
	"IthomeCralwer/cralwer/model"
	"IthomeCralwer/cralwer/tools"
	"regexp"
)

const titleRe = `<h1>(.*?)</h1>`
const timeRe = `<span id="pubtime_baidu">(.*?)</span>`

const articleListRe = `https://www.ithome.com/0/[\d]{3}/[\d]{3}.htm`

// ParseArticle here
func ParseArticle(contents []byte, _url string) engine.ParseResult {
	article := model.Article{}
	title := tools.ReSubMatch(titleRe, string(contents))
	article.Title = title
	time := tools.ReSubMatch(timeRe, string(contents))
	article.Time = time
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
