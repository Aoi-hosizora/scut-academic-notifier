package main

import (
	"fmt"
	"time"

	"gopkg.in/fatih/set.v0"
)

const NORMAL_NEWS_CNT int = 5

var TimeInternal time.Duration = 10 * time.Minute

var JWurl string = "http://jwc.scuteo.com/jiaowuchu/cms/index.do"
var JsonPath string = "./config.json"

func main() {

	var SCKEY string = getConfig(JsonPath).SCKEY
	if SCKEY == "" {
		panic("SCKEY is null")
	}

	defer func() {
		if err := recover(); err != nil {
			var msg string = fmt.Sprintf("> Panic: %s\n", err)
			putNotifier(SCKEY, "教务系统通知", msg)
			fmt.Println(msg)
		}
	}()

	fmt.Printf("\n> 开始监听，每 %s 获取数据一次...\n", TimeInternal)
	handleNewNotice(JWurl, SCKEY)
}

func handleNewNotice(url string, SCKEY string) {
	newSet := set.New(set.ThreadSafe)

	for {
		time.Sleep(TimeInternal)
		fmt.Println(getTimeString())

		tops, normals := handleData(getHTML(url))
		notice := getNewNews(tops, normals)

		diff := getSetStr(set.Difference(getSet(notice), newSet))
		if diff != "" {
			putNotifier(SCKEY, "教务系统通知", diff)
			fmt.Println(diff)
		}

		newSet = getSet(notice)

	}

}
