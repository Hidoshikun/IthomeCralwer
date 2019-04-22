package engine

// Request includes URL and ParserFunc
type Request struct {
	URL        string
	ParserFunc func([]byte, string) ParseResult
}

// ParseResult includes Requests and Items
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

// NilParser here
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
