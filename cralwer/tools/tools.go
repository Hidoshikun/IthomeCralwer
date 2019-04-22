package tools

import (
	"bufio"
	"log"
	"regexp"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

// ReMatch Apply regexp
func ReMatch(regex string, content string) string {
	re := regexp.MustCompile(regex)
	match := re.FindStringSubmatch(content)
	if len(match) == 0 {
		return ""
	}
	return match[0]
}

// ReSubMatch Apply regexp
func ReSubMatch(regex string, content string) string {
	re := regexp.MustCompile(regex)
	match := re.FindStringSubmatch(content)
	if len(match) == 0 {
		return ""
	}
	return match[1]
}

// DetermineEncoding parse the encoding for the given text
func DetermineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("DetermineEncoding error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
