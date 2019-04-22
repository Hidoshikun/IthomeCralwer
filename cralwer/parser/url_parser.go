package parser

import (
	"ithome/cralwer/engine"
	"ithome/cralwer/model"
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// const articleListRe = `https://www.ithome.com/(.*?).htm`

// ParseArticleList return article url list
/***
func ParseArticleList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(articleListRe)
	matches := re.FindAllString(string(contents), -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				URL:        string(m),
				ParserFunc: ParseArticle,
			})
	}
	return result
}
***/

const articleURLRe = `https://www.ithome.com/(.*?).htm`

const articleTimeRe = `[\d]{4}-[\d]{2}-[\d]{2} [\d]{2}:[\d]{2}:[\d]{2}`

// ParseArticleList return article url list
func ParseArticleList(contents []byte, _url string) engine.ParseResult {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		panic(err)
	}
	parserResult := engine.ParseResult{}

	dom.Find("#wrapper > div.content.fl > div > ul > li").
		Each(func(i int, selection *goquery.Selection) {
			content, err := selection.Html()
			if err != nil {
				log.Printf("test error %v", err)
			}
			articleURL := model.ArticleURL{
				URL:  ReMatch(articleURLRe, content),
				Time: ReMatch(articleTimeRe, content),
			}
			parserResult.Items = append(parserResult.Items, articleURL)
		})
	return parserResult
}

// ReMatch Apply regexp
func ReMatch(regex string, content string) string {
	re := regexp.MustCompile(regex)
	match := re.FindStringSubmatch(content)
	if len(match) == 0 {
		return ""
	}
	return match[0]
}