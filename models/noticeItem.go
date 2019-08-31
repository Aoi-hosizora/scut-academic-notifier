package models

import (
	"fmt"
)

type NoticeItem struct {
	Title string
	Url   string
	Type  string
	Date  string
	IsNew bool
}

func (i *NoticeItem) String() string {
	return fmt.Sprintf("%s：[%s](%s) - %s", i.Type, i.Title, i.Url, i.Date)
}
