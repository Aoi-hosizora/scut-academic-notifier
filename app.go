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

// 教务通知链接
var JWUrl string = "http://jw.scut.edu.cn/zhinan/cms/index.do"

// 教务通知 API
var JWAPIUrl string = "http://jw.scut.edu.cn/zhinan/cms/article/v2/findInformNotice.do"

// 软件学软 新闻资讯 链接
var SEUrl string = "http://www2.scut.edu.cn/sse/%s/list.htm"

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

var newSet = make([]models.NoticeItem, SendMaxCnt)

// 获取教务通知，判断更新
func grabNotice(url string, SCKEY string) {

	defer func() {
		if err := recover(); err != nil {
			var msg string = fmt.Sprintf("> Panic: %s\n\n+ 忽略 Panic 继续监听中...", err)
			utils.SendNotifier(SCKEY, "教务系统通知 错误信息", msg)
			log.Println(err)
		}
	}()

	for {
		// 通知
		notices := utils.ParseJson(utils.GetPostData(url, 0, 50))
		// 差集
		diffs := utils.ToArrayDifference(notices, newSet)

		// 向上取整
		ceil := int(math.Ceil(float64(len(diffs)) / float64(SendMaxCnt)))
		for i := 0; i < ceil; i++ {
			msg := ""
			for j := i * SendMaxCnt; j < i*SendMaxCnt+SendMaxCnt; j++ {
				if j < len(diffs) {
					ni := diffs[j]
					msg = msg + fmt.Sprintf("+ %s\n", ni.String())
				} else {
					break
				}
			}
			if msg != "" {
				if i == ceil-1 {
					msg += fmt.Sprintf("--- \n+ 更多通知请访问[华工教务通知](%s)", JWUrl)
				}
				title := ""
				if ceil == 1 {
					title = "教务系统通知"
				} else {
					title = fmt.Sprintf("教务系统通知_%d_%d", i+1, ceil)
				}
				utils.SendNotifier(SCKEY, title, msg)
				fmt.Println(msg)
			}
		}

		newSet = notices
		fmt.Println(utils.GetNowTimeString())
		time.Sleep(TimeInternal)
	}
}
