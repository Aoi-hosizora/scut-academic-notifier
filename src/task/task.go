package task

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xslice"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/bot/server"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/config"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/database"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/model"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/service"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/static"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/wechat"
	"github.com/robfig/cron/v3"
	"gopkg.in/tucnak/telebot.v2"
	"math"
	"strings"
	"sync"
)

var Cron *cron.Cron

func Setup() error {
	Cron = cron.New(cron.WithSeconds())

	_, err := Cron.AddFunc(config.Configs.Task.Cron, task)
	if err != nil {
		return err
	}

	return nil
}

func task() {
	defer func() { recover() }()

	users := database.GetUsers()
	if len(users) == 0 {
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(len(users))
	for _, user := range users {
		go func(user *model.User) {
			doForUser(user)
			wg.Done()
		}(user)
	}
	wg.Wait()
}

func doForUser(user *model.User) {
	// get new data
	newJwItems, err := service.GetJwItems()
	if err != nil || len(newJwItems) == 0 {
		return
	}
	newSeItems, err := service.GetSeItems()
	if err != nil || len(newSeItems) == 0 {
		return
	}

	// filter new data
	newItems := make([]*model.Item, 0)
	for _, jw := range newJwItems {
		if service.CheckTime(jw.Date, config.Configs.Send.Range) {
			newItems = append(newItems, jw)
		}
	}
	for _, se := range newSeItems {
		if service.CheckTime(se.Date, config.Configs.Send.Range) {
			newItems = append(newItems, se)
		}
	}

	// get old data and calc diff
	oldItems, ok := database.GetOldData(user.ChatID)
	if !ok {
		return
	}
	sendItems := xslice.Its(xslice.DiffWith(xslice.Sti(newItems), xslice.Sti(oldItems), func(i, j interface{}) bool {
		i1 := i.(*model.Item)
		i2 := j.(*model.Item)
		return i1.Type == i2.Type && i1.Title == i2.Title
	}), &model.Item{}).([]*model.Item)
	if len(sendItems) == 0 {
		return
	}

	// send items
	moreStr := fmt.Sprintf("更多信息，请查阅 [华工教务通知](%s) 以及 [软院公务通知](%s)。", static.JwHomepage, static.SeHomepage)

	// send bot
	sb := strings.Builder{}
	sb.WriteString("*学校相关通知*\n=====\n")
	for idx, item := range sendItems {
		sb.WriteString(fmt.Sprintf("%d. %s\n", idx+1, item.String()))
	}
	sb.WriteString("=====\n")
	sb.WriteString(moreStr)

	msg := sb.String()
	err = server.Bot.SendToChat(user.ChatID, msg, telebot.ModeMarkdown)
	if err != nil {
		return
	}

	// send wechat
	maxCnt := int(config.Configs.Send.MaxCount)
	sendTimes := int(math.Ceil(float64(len(sendItems)) / float64(maxCnt)))
	for i := 0; i < sendTimes; i++ {
		from := i * maxCnt
		to := (i + 1) * maxCnt
		if l := len(sendItems); to > l {
			to = l
		}

		sb := strings.Builder{}
		for j := from; j < to; j++ {
			sb.WriteString(fmt.Sprintf("%d. %s\n", j+1, sendItems[j].String()))
		}
		if i == sendTimes-1 {
			sb.WriteString("\n--- \n")
			sb.WriteString(moreStr)
		}

		msg := sb.String()
		title := fmt.Sprintf("学校相关通知 (第 %d 条，共 %d 条)", i+1, sendTimes)
		_ = wechat.SendToChat(user.Sckey, title, msg)
	}

	// write old data
	database.SetOldData(user.ChatID, newItems)
}
