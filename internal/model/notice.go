package model

import (
	"github.com/Aoi-hosizora/ahlib/xgeneric/xgslice"
	"sort"
)

type NoticeItem struct {
	Title string `json:"title"` // <- will be saved to redis
	Url   string `json:"url"`
	Type  string `json:"type"` // <- will be saved to redis
	Date  string `json:"date"`
}

type NoticeItemResult struct {
	Data struct {
		Data []*NoticeItem `json:"data"`
	} `json:"data"`
}

func DiffNoticeItemSlice(s1 []*NoticeItem, s2 []*NoticeItem) []*NoticeItem {
	return xgslice.DiffWith(s1, s2, func(i, j *NoticeItem) bool {
		return i.Title == j.Title && i.Type == j.Type
	})
}

func SortNoticeItemSlice(items []*NoticeItem) {
	sort.Slice(items, func(i, j int) bool {
		return items[i].Date > items[j].Date // reverse
	})
}
