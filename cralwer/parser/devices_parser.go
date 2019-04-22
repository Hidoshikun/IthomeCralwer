package parser

import (
	"IthomeCralwer/cralwer/engine"
	"IthomeCralwer/cralwer/model"
	"IthomeCralwer/cralwer/tools"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/bitly/go-simplejson"
)

const commentRe = `href="//m.ithome.com/ithome/download/">(.*?)</a></span><span class="posandtime">`

const aritlcleIDRe = `NewsID=([\d]*)&LapinID`

// ParseCommentDevices return article url list
func ParseCommentDevices(contents []byte, url string) engine.ParseResult {
	if len(contents) <= 13 {
		// 返回{"Success":0}
		return engine.ParseResult{}
	}

	res, err := simplejson.NewJson(contents)
	if err != nil {
		fmt.Printf("err was %v", err)
	}

	rows, err := res.Get("Result").Get("Clist").Array()
	if err != nil {
		fmt.Printf("json parse error was:  %v", err)
	}

	if err != nil {
		fmt.Printf("json parse error was:  %v", err)
	}

	deviceList := []string{}

	minCommentID := 100000000

	for _, row := range rows {
		if eachMap, ok := row.(map[string]interface{}); ok {
			item := eachMap["M"].(map[string]interface{})
			deviceList = append(deviceList, item["Ta"].(string))
			commentID, err := strconv.Atoi(string(item["Ci"].(json.Number)))
			if err != nil {
				fmt.Printf("strconv.Atoi(item['Ta'].(string)) error was:  %v", err)
			}

			if commentID <= minCommentID {
				minCommentID = commentID
			}
			replys := eachMap["R"].([]interface{})
			// 提取楼中楼
			for _, reply := range replys {
				eachR := reply.(map[string]interface{})
				deviceList = append(deviceList, eachR["Ta"].(string))
			}
		}
	}

	parserResult := engine.ParseResult{}
	articleID, err := strconv.Atoi(tools.ReSubMatch(aritlcleIDRe, url))
	if err != nil {
		fmt.Printf("json parse error was:  %v", err)
	}

	// 构造item
	item := model.DeviceList{
		ArticleID:  articleID,
		DeviceList: deviceList,
	}
	request := engine.Request{
		URL:        "https://m.ithome.com/api/comment/newscommentlistget?NewsID=" + strconv.Itoa(articleID) + "&LapinID=&MaxCommentID=" + strconv.Itoa(minCommentID) + "&Latest=",
		ParserFunc: ParseCommentDevices,
	}
	parserResult.Items = append(parserResult.Items, item)
	parserResult.Requests = append(parserResult.Requests, request)

	return parserResult
}
