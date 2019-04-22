package fetcher

import (
	"io/ioutil"
	"ithome/cralwer/tools"
	"net/http"
)

const postURL = "https://dyn.ithome.com/ithome/getajaxdata.aspx"

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
	hash := tools.ReSubMatch(hashRe, string(content))
	return hash, nil
}
