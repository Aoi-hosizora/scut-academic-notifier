package main

import "fmt"

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
	str := fmt.Sprintf("%s%s (%s)", isnew, news.title, news.href)
	return str
}

func (news *SCUT_NormalNews) String() string {
	isnew := ""
	if news.isnew {
		isnew = "[NEW]"
	}
	str := fmt.Sprintf("%s%s - %s (%s)", isnew, news.title, news.date, news.href)
	return str
}
