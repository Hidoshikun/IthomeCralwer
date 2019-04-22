package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

const postURL = "https://dyn.ithome.com/ithome/getajaxdata.aspx"

const articleIDRe = `/([\d]*).htm`

// POST请求旧版本
/***
// CommentFetcher content from given url
func CommentFetcher(urla string) ([]byte, error) {
	fmt.Println("fetching url: ", urla)
	articleID := ReSubMatch(articleIDRe, urla)
	urlb := "https://dyn.ithome.com/comment/" + articleID
	hash, err := getHash(urlb)
	if err != nil {
		return nil, err
	}
	content := []byte{}
	// 构建POST请求，循环翻页
	for i := 1; ; i++ {
		resp, err := http.PostForm(postURL,
			url.Values{
				"newsID": {articleID},
				"hash":   {hash},
				"type":   {"commentpage"},
				"page":   {strconv.Itoa(i)},
				"order":  {"false"}})

		if err != nil {
			log.Printf("post error happend: %v", err)
			return nil, err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("read body error happend: %v", err)
			return nil, err
		}
		content = append(content, body...)
		// 翻页为空时退出循环
		if len(body) == 0 {
			break
		}
	}

	return content, nil
}
***/

// CommentFetcher content from given url
func CommentFetcher(urla string) ([]byte, error) {
	fmt.Println("fetching url: ", urla)
	articleID := ReSubMatch(articleIDRe, urla)
	urlb := "https://m.ithome.com/api/comment/newscommentlistget?NewsID=" + articleID + "&LapinID=&MaxCommentID=0&Latest="
	resp, err := http.Get(urlb)
	if err != nil {
		log.Printf("get comment error happend: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code : %d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

/***
	urlc := "https://dyn.ithome.com/ithome/getajaxdata.aspx"
	resp, err := http.PostForm(urlc,
		url.Values{"key": {"Value"}, "id": {"123"}})
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code : %d", resp.StatusCode)
	}
	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}
***/

// 获取文章Hash值
func getHash(urla string) (string, error) {
	resp, err := http.Get(urla)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	hashRe := "var ch11 = '(.*?)';</script>"
	hash := ReSubMatch(hashRe, string(content))
	return hash, nil
}

// ReMatch get the complite string
func ReMatch(regex string, content string) string {
	re := regexp.MustCompile(regex)
	match := re.FindString(content)
	return match
}

// ReSubMatch get the sub string
func ReSubMatch(regex string, content string) string {
	re := regexp.MustCompile(regex)
	match := re.FindStringSubmatch(content)
	if len(match) == 0 {
		return ""
	}
	return match[1]
}
