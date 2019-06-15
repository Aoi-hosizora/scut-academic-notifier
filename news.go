package main

import "fmt"

type SCUT_TopNews struct {
	title string
	href  string
}

type SCUT_NormalNews struct {
	title string
	href  string
	date  string
}

// type SCUT_NEWS struct {
// 	top    SCUT_TopNews
// 	normal SCUT_NormalNews
// }

func (news *SCUT_TopNews) String() string {
	str := fmt.Sprintf("%s (%s)", news.title, news.href)
	return str
}

func (news *SCUT_NormalNews) String() string {
	str := fmt.Sprintf("%s - %s (%s)", news.title, news.date, news.href)
	return str
}
