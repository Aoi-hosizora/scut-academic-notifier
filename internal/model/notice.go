package model

import (
	"github.com/Aoi-hosizora/ahlib/xslice"
)

type NoticeItem struct {
	Title string `json:"title"` // <- will be saved in redis
	Url   string `json:"url"`
	Type  string `json:"type"` // <- will be saved in redis
	Date  string `json:"date"`
}

type NoticeItemResult struct {
	Data struct {
		Data []*NoticeItem `json:"data"`
	} `json:"data"`
}

func DiffNoticeItemSlice(s1 []*NoticeItem, s2 []*NoticeItem) []*NoticeItem {
	return xslice.DiffWithG(s1, s2, func(i, j interface{}) bool {
		p1, p2 := i.(*NoticeItem), j.(*NoticeItem)
		return p1.Title == p2.Title && p1.Type == p2.Type
	}).([]*NoticeItem)
}
