package main

import (
	"container/list"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type pList = *list.List
type ListArrNoTop = [NORMAL_NEWS_CNT](pList)
type ListArrAndTop = [NORMAL_NEWS_CNT + 1](pList)

type pSel = *goquery.Selection
type pDoc = *goquery.Document

type SCUT_TopNews struct {
	title string
	href  string
	isnew bool
}

type SCUT_NormalNews struct {
	title string
	href  string
	date  string
	isnew bool
}

func (news *SCUT_TopNews) String() string {
	isnew := ""
	if news.isnew {
		isnew = "[NEW]"
	}
	str := fmt.Sprintf("%s%s", isnew, news.title)
	return str
}

func (news *SCUT_NormalNews) String() string {
	isnew := ""
	if news.isnew {
		isnew = "[NEW]"
	}
	str := fmt.Sprintf("%s%s - %s)", isnew, news.title, news.date)
	return str
}
