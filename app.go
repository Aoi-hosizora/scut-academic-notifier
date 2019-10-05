package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/Aoi-hosizora/Academic_Notifier/models"
	"github.com/Aoi-hosizora/Academic_Notifier/utils"
)

// 访问教务通知频率
var TimeInternal time.Duration = 10 * time.Minute

// 一次发送的最大量
var SendMaxCnt int = 10

// 连续错误最大次数
var ErrorMaxCnt int = 5

// 发送最久半个月前的信息
var SendTimeRange time.Duration = 30 * 24 * time.Hour

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
var oldSeSet = make([]models.NoticeItem, 50)

// 记录连续错误次数
var errCnt = 0

// 获取教务通知，判断更新
func grabNotice(url string, SCKEY string) {

	moreStr := fmt.Sprintf("--- \n+ 更多通知请访问[华工教务通知](%s)和[软院公务通知](%s)", JWUrl, fmt.Sprintf(SEUrl, SEUrlParts[0]))

	defer func() {
		if err := recover(); err != nil {
			errCnt += 1
			msg := ""
			if errCnt < ErrorMaxCnt {
				msg = fmt.Sprintf("+ 忽略 Panic (第 %d 次) 继续监听中...\n\n+ Panic: %s\n"+moreStr, errCnt, err)
			} else {
				msg = fmt.Sprintf("+ Panic 已经超过 %d 次，程序终止\n\n+ Panic: %s\n"+moreStr, errCnt, err)
			}
			if len(msg) >= 203 {
				msg = msg[0:200] + "..."
			}
			utils.SendNotifier(SCKEY, "教务系统通知 错误信息", msg)
			log.Println(err)
			log.Printf("errCnt: %d", errCnt)
			if errCnt >= ErrorMaxCnt {
				panic("")
			}
		}
	}()

	for {
		// 通知
		newSet := utils.ParseJson(utils.GetPostData(url, 0, 50))
		// newSet := []models.NoticeItem{}
		newSeSet := utils.GetSENotices(SEUrl, SEUrlParts, SEUrlPartNames)
		if newSeSet != nil {
			oldSeSet = newSeSet
		}

		// 并集差集
		newSet = utils.ToArrayAdd(newSet, oldSeSet)
		diffs := utils.ToArrayDifference(newSet, oldSet)

		// 一个月内的
		sendLists := make([]models.NoticeItem, 0)
		for _, v := range diffs {
			nt, err := time.ParseInLocation("2006-01-02 15:04:05", v.Date+" 00:00:00", time.Local)
			if err == nil && nt.After(time.Now().Add(-SendTimeRange)) {
				sendLists = append(sendLists, v)
			}
		}

		// 向上取整
		ceil := int(math.Ceil(float64(len(sendLists)) / float64(SendMaxCnt)))
		for i := 0; i < ceil; i++ {
			// 切分合并记录
			msg := ""
			for j := i * SendMaxCnt; j < i*SendMaxCnt+SendMaxCnt; j++ {
				if j < len(sendLists) {
					ni := sendLists[j]
					msg += fmt.Sprintf("+ %s\n", ni.String())
				} else {
					break
				}
			}
			// 发送
			if msg != "" {
				// 最后一条
				if i == ceil-1 {
					msg += moreStr
				}
				utils.SendNotifier(SCKEY, fmt.Sprintf("教务系统通知（第 %d 条，共 %d 条）", i+1, ceil), msg)
				fmt.Println("已发送：\n" + msg)
			}
		}

		oldSet = newSet
		fmt.Println(utils.GetNowTimeString())
		errCnt = 0
		time.Sleep(TimeInternal)
	}
}
