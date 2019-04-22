package parser

import (
	"IthomeCralwer/cralwer/engine"
	"IthomeCralwer/cralwer/model"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/bitly/go-simplejson"
)

// ParseComment here
func ParseComment(contents []byte, url string) engine.ParseResult {
	res, err := simplejson.NewJson(contents)
	if err != nil {
		fmt.Printf("err was %v", err)
	}

	rows, err := res.Get("Result").Get("Clist").Array()
	if err != nil {
		fmt.Printf("err was %v", err)
	}

	parserResult := engine.ParseResult{}

	for _, row := range rows {
		if eachMap, ok := row.(map[string]interface{}); ok {
			// fmt.Println("M is: ", eachMap["M"])
			item := eachMap["M"].(map[string]interface{})
			userID, err := strconv.Atoi(string(item["Ci"].(json.Number)))
			if err != nil {
				panic(err)
			}

			supported, err := strconv.Atoi(string(item["S"].(json.Number)))
			if err != nil {
				panic(err)
			}

			aganist, err := strconv.Atoi(string(item["A"].(json.Number)))
			if err != nil {
				panic(err)
			}

			comment := model.Comment{
				UserID:     userID,
				Content:    DelEmoji(item["C"].(string)),
				UserName:   item["N"].(string),
				Location:   item["Y"].(string),
				Time:       item["T"].(string),
				Device:     item["Ta"].(string),
				Supported:  supported,
				Aganist:    aganist,
				ArticleURL: url,
			}
			parserResult.Items = append(parserResult.Items, comment)
		}
	}

	return parserResult
}

// DelEmoji here
func DelEmoji(s string) string {
	ret := ""
	rs := []rune(s)
	for i := 0; i < len(rs); i++ {
		if len(string(rs[i])) == 4 {
			u := `[\u` + strconv.FormatInt(int64(rs[i]), 16) + `]`
			ret += u

		} else {
			ret += string(rs[i])
		}
	}
	return ret
}
