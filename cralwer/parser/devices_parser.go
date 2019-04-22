package parser

import (
	"fmt"
	"ithome/cralwer/engine"
	"ithome/cralwer/model"

	"github.com/bitly/go-simplejson"
)

const commentRe = `href="//m.ithome.com/ithome/download/">(.*?)</a></span><span class="posandtime">`

// ParseCommentDevices return article url list
func ParseCommentDevices(contents []byte, url string) engine.ParseResult {
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
	for _, row := range rows {
		if eachMap, ok := row.(map[string]interface{}); ok {
			// fmt.Println("M is: ", eachMap["M"])
			item := eachMap["M"].(map[string]interface{})
			deviceList = append(deviceList, item["Ta"].(string))
		}
	}

	parserResult := engine.ParseResult{}
	// 构造item
	item := model.DeviceList{
		URL:        url,
		DeviceList: deviceList,
	}
	parserResult.Items = append(parserResult.Items, item)

	return parserResult
}
