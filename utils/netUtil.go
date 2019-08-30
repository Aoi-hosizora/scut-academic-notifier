package utils

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// Server Chan Url: `Sckey` `title` `msg`
var ServerChanUrl string = "https://sc.ftqq.com/%s.send?text=%s&desp=%s"

type pSel = *goquery.Selection
type pDoc = *goquery.Document
type pList = *list.List

// 通过 Server 酱发送信息
func SendNotifier(Sckey string, title string, msg string) {

	// 将发送内容加上时间
	msg = fmt.Sprintf("> At %s Send: \n %s", GetNowTimeString(), msg)

	// url.QueryEscape 转化 url
	url := fmt.Sprintf(ServerChanUrl, Sckey, url.QueryEscape(title), url.QueryEscape(msg))
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	if res.StatusCode != 200 {
		panic(res.Status)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
}

// Html Get Doc
func GetHTMLDoc(url string) pDoc {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != 200 {
		panic(res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}
	return doc
}

// 解析 教务通知 Doc
func ParseData(doc pDoc) string {
	// rets := list.New()
	str, _ := doc.Find(".posts_right").Html()
	return str
	// return rets
}
