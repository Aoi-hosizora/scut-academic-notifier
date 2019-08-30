package models

import "fmt"

type NoticeItem struct {
	Title string
	Url   string
	Type  string
	Date  string
	IsNew bool
}

func (i *NoticeItem) String() string {
	isNew := ""
	if i.IsNew {
		isNew = "[New] "
	}
	return fmt.Sprintf("%s%s: [%s](%s) - %s)", isNew, i.Type, i.Title, i.Url, i.Date)
}
