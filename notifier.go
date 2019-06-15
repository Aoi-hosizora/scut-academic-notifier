package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const NORMAL_NEWS_CNT int = 5

var JsonPath string = "./config.json"

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("> Panic: %s\n", err)
		}
	}()

	var SCKEY string = getConfig(JsonPath).SCKEY
	if SCKEY == "" {
		panic("SCKEY is null")
	}

	var url string = "http://jwc.scuteo.com/jiaowuchu/cms/index.do"

	var doc *goquery.Document = getHTML(url)
	tops, normals := handleData(doc)
	printAllData(tops, normals)

	putNotifier(SCKEY, "Test", "Test")
}

func printAllData(tops *list.List, normals [NORMAL_NEWS_CNT](*list.List)) {

	fmt.Printf("\n////// 顶部信息 //////\n\n")
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

type SckeyConfig struct {
	SCKEY string
}

func getConfig(url string) *SckeyConfig {
	cfg := &SckeyConfig{}
	data, err := ioutil.ReadFile(url)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

func putNotifier(Sckey string, title string, msg string) {

	msg = fmt.Sprintf("%s (%s)", msg, time.Now().Format("2006-01-02 15:04:05"))

	// url.QueryEscape 转化 url
	url := fmt.Sprintf("https://sc.ftqq.com/%s.send?text=%s&desp=%s", Sckey, url.QueryEscape(title), url.QueryEscape(msg))

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

	print(string(body))
}
