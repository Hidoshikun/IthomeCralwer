package engine

import (
	"ithome/cralwer/fetcher"
	"log"
)

func worker(r Request) (ParseResult, error) {
	body, err := fetcher.CommentFetcher(r.URL)
	if err != nil {
		log.Printf("Fetch: error fetching url %s: %v", r.URL, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body, r.URL), nil
}
