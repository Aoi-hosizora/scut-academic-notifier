package main

import (
	"fmt"
	"log"

	"time"

	"github.com/Aoi-hosizora/Academic_Notifier/models"
	"github.com/Aoi-hosizora/Academic_Notifier/utils"
)

// 访问教务通知频率
var TimeInternal time.Duration = 10 * time.Minute

// 一次发送的最大量
var SendMaxCnt int = 50

// 发送最久半个月前的信息
var SendTimeRange time.Duration = 15 * 24 * time.Hour

// 教务通知链接
var JWUrl string = "http://jw.scut.edu.cn/zhinan/cms/index.do"

// 教务通知 API
var JWAPIUrl string = "http://jw.scut.edu.cn/zhinan/cms/article/v2/findInformNotice.do"

// 软件学软 新闻资讯 链接
var SEUrl string = "http://www2.scut.edu.cn/sse/%s/list.htm"
var SEUrlParts = []string{"xyjd_17232", "17235", "17236", "gwtz", "kytz"}
var SEUrlPartNames = []string{"学院焦点", "本科生通知", "研究生通知", "公务通知", "科研通知"}

// 配置文件路径
var JsonPath string = "./config.json"

func main() {

	var SCKEY string = utils.GetConfig(JsonPath).SCKEY
	if SCKEY == "" {
		panic("SCKEY is null")
	}

	fmt.Printf("\n> 开始监听，每 %s 获取数据一次...\n", TimeInternal)

	for {
		grabNotice(JWAPIUrl, SCKEY)
	}
}

var oldSet = make([]models.NoticeItem, SendMaxCnt)
var oldSeSet = make([]models.NoticeItem, SendMaxCnt)

// 获取教务通知，判断更新
func grabNotice(url string, SCKEY string) {

	moreStr := fmt.Sprintf("--- \n+ 更多通知请访问[华工教务通知](%s)", JWUrl)

	defer func() {
		if err := recover(); err != nil {
			var msg string = fmt.Sprintf("> Panic: %s\n\n+ 忽略 Panic 继续监听中...\n"+moreStr, err)
			utils.SendNotifier(SCKEY, "教务系统通知 错误信息", msg)
			log.Println(err)
		}
	}()

	for {
		// 通知
		// newSet := utils.ParseJson(utils.GetPostData(url, 0, 50))
		newSet := []models.NoticeItem{}
		newSeSet := utils.GetSENotices(SEUrl, SEUrlParts, SEUrlPartNames)
		if newSeSet != nil {
			oldSeSet = newSeSet
		}
		newSet = utils.ToArrayAdd(newSet, oldSeSet)

		// 差集
		diffs := utils.ToArrayDifference(newSet, oldSet)

		// 信息
		msg := ""
		for _, v := range diffs {

			// 一个月内的
			nt, err := time.ParseInLocation("2006-01-02 15:04:05",
				v.Date+" 00:00:00", time.Local)

			if err == nil {
				if nt.Before(time.Now().Add(-SendTimeRange)) {
					continue
				}
			}
			msg = msg + fmt.Sprintf("+ %s\n", v.String())
		}

		// 发送
		if msg != "" {
			msg += moreStr
			utils.SendNotifier(SCKEY, "教务系统通知", msg)
			fmt.Println("已发送：\n" + msg)
		}

		oldSet = newSet
		fmt.Println(utils.GetNowTimeString())
		time.Sleep(TimeInternal)
	}
}
