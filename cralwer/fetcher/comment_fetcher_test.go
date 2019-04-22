package fetcher

import (
	"fmt"
	"testing"

	simplejson "github.com/bitly/go-simplejson"
)

func TestPostFetcher(t *testing.T) {
	urla := "https://www.ithome.com/html/it/340608.htm"
	html, err := CommentFetcher(urla)
	if err != nil {
		panic(err)
	}

	res, err := simplejson.NewJson(html)
	if err != nil {
		fmt.Printf("err was %v", err)
	}

	rows, err := res.Get("Result").Get("Clist").Array()

	for _, row := range rows {
		if eachMap, ok := row.(map[string]interface{}); ok {
			// fmt.Println("M is: ", eachMap["M"])
			// item := eachMap["M"].(map[string]interface{})
			replys := eachMap["R"].([]interface{})
			// fmt.Println(item["Ci"], item["N"], item["Y"], item["T"], item["Ta"])
			for _, reply := range replys {
				fmt.Println("reply is : ", reply)
				eachR := reply.(map[string]interface{})
				fmt.Println(eachR["Ci"], eachR["N"], eachR["Y"], eachR["T"], eachR["Ta"])
			}
		}
	}
	fmt.Println("row len: ", len(rows))
}
