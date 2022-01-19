package model

import (
	"github.com/Aoi-hosizora/ahlib/xslice"
)

type PostItem struct {
	Title string `json:"title"`
	Url   string `json:"url"`
	Type  string `json:"type"`
	Date  string `json:"date"`
}

type PostItemDto struct {
	Data struct {
		Data []*PostItem `json:"data"`
	} `json:"data"`
}

func ItemSliceDiff(s1 []*PostItem, s2 []*PostItem) []*PostItem {
	return xslice.DiffWithG(s1, s2, func(i, j interface{}) bool {
		p1, p2 := i.(*PostItem), j.(*PostItem)
		return p1.Title == p2.Title && p1.Type == p2.Type
	}).([]*PostItem)
}
