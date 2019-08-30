package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Aoi-hosizora/Academic_Notifier/utils"
)

// 访问教务通知频率
var TimeInternal time.Duration = 10 * time.Minute

// 教务通知链接
var JWUrl string = "http://jw.scut.edu.cn/zhinan/cms/toPosts.do"

// 配置文件路径
var JsonPath string = "./config.json"

func main() {

	var SCKEY string = utils.GetConfig(JsonPath).SCKEY
	if SCKEY == "" {
		panic("SCKEY is null")
	}

	defer func() {
		if err := recover(); err != nil {
			var msg string = fmt.Sprintf("> Panic: %s\n", err)
			utils.SendNotifier(SCKEY, "教务系统通知", msg)
			log.Fatal(msg)
		}
	}()

	fmt.Printf("\n> 开始监听，每 %s 获取数据一次...\n", TimeInternal)
	grabNotice(JWUrl, SCKEY)
}

// 获取教务通知，判断更新
func grabNotice(url string, SCKEY string) {
	// newSet := set.New(set.ThreadSafe)

	// for {
	// 	time.Sleep(TimeInternal)
	// 	fmt.Println(utils.GetNowTimeString())

	// 	tops, normals := handleData(getHTML(url))
	// 	notice := getNewNews(tops, normals)

	// 	diff := getSetStr(set.Difference(getSet(notice), newSet))
	// 	if diff != "" {
	// 		putNotifier(SCKEY, "教务系统通知", diff)
	// 		fmt.Println(diff)
	// 	}
	// 	newSet = getSet(notice)
	// }

	fmt.Println(utils.ParseData(utils.GetHTMLDoc(url)))
}
