package models

import (
	"fmt"
	"strconv"
	"strings"
)

type NoticeItem struct {
	Title string
	Url   string
	Type  string
	Date  string
	IsNew bool
}

func (i *NoticeItem) String() string {
	// isNew := ""
	// if i.IsNew {
	// 	isNew = "[New] "
	// }
	// return fmt.Sprintf("%s%s: [%s](%s) - %s)", isNew, i.Type, i.Title, i.Url, i.Date)
	return fmt.Sprintf("+ %s: [%s](%s) - %s", i.Type, i.Title, i.Url, i.Date)
}

type NoticeItems []NoticeItem

func (n NoticeItems) Len() int {
	return len(n)
}

func (n NoticeItems) Less(i, j int) bool {
	n1, err1 := strconv.Atoi(strings.Replace(n[i].Date, ".", "", -1))
	n2, err2 := strconv.Atoi(strings.Replace(n[j].Date, ".", "", -1))
	if err1 != nil || err2 != nil {
		return false
	}
	return n1 > n2
}

func (n NoticeItems) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
