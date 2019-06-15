package main

import (
	"container/list"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

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

		// 栏目顶部信息
		icit.Find(".index-content-infos-tab-ad-info").Each(func(i int, s *goquery.Selection) {

			isnew := false
			s.Find("div>img").Each(func(i int, s *goquery.Selection) {
				isnew = true
			})

			top := &SCUT_TopNews{
				title: s.Find("a").Text(),
				href:  s.Find("a").AttrOr("href", "."),
				isnew: isnew,
			}
			tops.PushBack(top)
		})

		// 栏目列表信息
		normals[normal_idx] = list.New()
		icit.Find(".index-content-infos-tab-ad-news > li").Each(func(i int, s *goquery.Selection) {
			isnew := false
			s.Find("div>img").Each(func(i int, s *goquery.Selection) {
				isnew = true
			})
			normal := &SCUT_NormalNews{
				title: s.Find("a").Text(),
				date:  s.Find("span").Text(),
				href:  s.Find("a").AttrOr("href", "."),
				isnew: isnew,
			}
			normals[normal_idx].PushBack(normal)
		})
		normal_idx = normal_idx + 1
	})

	return tops, normals
}
