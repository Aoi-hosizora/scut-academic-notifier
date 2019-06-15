package main

import (
	"container/list"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func getHTML(url string) pDoc {
	// 获取 HTML
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

func handleData(doc pDoc) (pList, ListArrNoTop) {
	// 获取所有提醒数据

	tops := list.New()
	var normals ListArrNoTop

	var normal_idx int = 0

	doc.Find(".index-content-infos > .index-content-infos-tab").Each(func(i int, s pSel) {

		var icit pSel = s

		// 栏目顶部信息
		icit.Find(".index-content-infos-tab-ad-info").Each(func(i int, s pSel) {

			isnew := false
			s.Find("div>img").Each(func(i int, s pSel) {
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
		icit.Find(".index-content-infos-tab-ad-news > li").Each(func(i int, s pSel) {
			isnew := false
			s.Find("div>img").Each(func(i int, s pSel) {
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

// "顶部信息", "教务通知", "交换交流", "新闻动态", "学院信息", "媒体关注"
func getNewNews(tops pList, normals ListArrNoTop) ListArrAndTop {
	// 获取标记为 New 的数据
	var retList ListArrAndTop

	newlist := list.New()
	for i := tops.Front(); i != nil; i = i.Next() {
		news := i.Value.(*SCUT_TopNews)
		if news.isnew {
			newlist.PushBack(news)
		}
	}
	retList[0] = newlist

	for i := 0; i < NORMAL_NEWS_CNT; i++ {
		newlist = list.New()
		for j := normals[i].Front(); j != nil; j = j.Next() {
			news := j.Value.(*SCUT_NormalNews)
			if news.isnew {
				newlist.PushBack(news)
			}
		}
		retList[i+1] = newlist
	}

	return retList
}
