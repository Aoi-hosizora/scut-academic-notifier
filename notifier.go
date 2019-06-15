package main

import (
	"container/list"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const NORMAL_NEWS_CNT int = 5

func main() {

	var url string = "http://jwc.scuteo.com/jiaowuchu/cms/index.do"

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var doc *goquery.Document = getHTML(url)

	tops, normals := handleData(doc)

	for i := tops.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	strs := []string{
		"教务通知", "交换交流", "新闻动态", "学院信息", "媒体关注",
	}
	for i := 0; i < NORMAL_NEWS_CNT; i++ {
		fmt.Printf("\n////// %s //////\n\n", strs[i])
		for j := normals[i].Front(); j != nil; j = j.Next() {
			fmt.Println(j.Value)
		}
	}

}

func getHTML(url string) *goquery.Document {

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != 200 {
		panic(res.Status)
	}

	// html, err := ioutil.ReadAll(res.Body)
	// fmt.Print(string(html))

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	return doc
}

func handleData(doc *goquery.Document) (*list.List, [NORMAL_NEWS_CNT](*list.List)) {

	tops := list.New()

	var normals [NORMAL_NEWS_CNT](*list.List)

	var normal_idx int = 0

	doc.Find(".index-content-infos > .index-content-infos-tab").Each(func(i int, s *goquery.Selection) {

		var icit *goquery.Selection = s
		icit.Find(".index-content-infos-tab-ad-info").Each(func(i int, s *goquery.Selection) {

			top := &SCUT_TopNews{
				title: s.Find("a").Text(),
				href:  s.Find("a").AttrOr("href", "."),
			}
			tops.PushBack(top)
		})

		normals[normal_idx] = list.New()

		icit.Find(".index-content-infos-tab-ad-news > li").Each(func(i int, s *goquery.Selection) {
			normal := &SCUT_NormalNews{
				title: s.Find("a").Text(),
				date:  s.Find("span").Text(),
				href:  s.Find("a").AttrOr("href", "."),
			}
			normals[normal_idx].PushBack(normal)
		})

		normal_idx = normal_idx + 1
	})

	return tops, normals
}
