package model

import "encoding/json"

//Article here
type Article struct {
	URL     string
	ID      string
	Title   string
	Time    string
	Content string
}

// ArticleURL here
type ArticleURL struct {
	URL  string
	Time string
}

// DeviceList here
type DeviceList struct {
	URL        string
	DeviceList []string
}

// Comment here
type Comment struct {
	UserID     int
	Content    string
	UserName   string
	Location   string
	Time       string
	Device     string
	Supported  int
	Aganist    int
	ArticleURL string
}

// FromJSONObj here
func FromJSONObj(o interface{}) (Article, error) {
	var article Article
	s, err := json.Marshal(o)
	if err != nil {
		return article, err
	}

	err = json.Unmarshal(s, &article)
	return article, err
}
