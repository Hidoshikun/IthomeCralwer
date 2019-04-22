package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/text/transform"
)

// Fetch content from given url
func Fetch(url string) ([]byte, error) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("fetching url: ", url)
	resp, err := http.Get(url)
	if err != nil {
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
