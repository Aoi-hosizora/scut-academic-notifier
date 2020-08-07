package model

import "fmt"

type Item struct {
	Title string
	Url   string
	Type  string
	Date  string
}

func (d *Item) String() string {
	return fmt.Sprintf("%s: [%s](%s) - %s", d.Type, d.Title, d.Url, d.Date)
}

// An item of jw.
type JwItem struct {
	List []*struct {
		CreateTime string `json:"createTime"`
		Id         string `json:"id"`
		IsNew      bool   `json:"isNew"`
		Tag        int    `json:"tag"`
		Title      string `json:"title"`
	} `json:"list"`
	Message string `json:"message"`
	PageNum int    `json:"pagenum"`
	Row     int    `json:"row"`
	Success bool   `json:"success"`
	Total   int    `json:"total"`
}

// The whole page of a tag.
type SePage struct {
	TagPart string
	TagName string
	Content string
}
