package task

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xslice"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/bot"
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

var (
	oldJwData = make(map[int64][]*model.Item)
	oldSeData = make(map[int64][]*model.Item)
)

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

	// get old data
	oldJwItems, ok := oldJwData[user.ChatID]
	if !ok || oldJwItems == nil {
		oldJwItems = make([]*model.Item, 0)
		oldJwData[user.ChatID] = oldJwItems
	}
	oldSeItems, ok := oldSeData[user.ChatID]
	if !ok || oldSeItems == nil {
		oldSeItems = make([]*model.Item, 0)
		oldSeData[user.ChatID] = oldSeItems
	}

	// calc diff
	sendJwItems := xslice.Its(xslice.Diff(xslice.Sti(newJwItems), xslice.Sti(oldJwItems)), &model.Item{}).([]*model.Item)
	sendSeItems := xslice.Its(xslice.Diff(xslice.Sti(newSeItems), xslice.Sti(oldSeItems)), &model.Item{}).([]*model.Item)
	if len(sendJwItems) == 0 || len(sendSeItems) == 0 {
		return
	}

	// filter valid
	sendItems := make([]*model.Item, 0)
	for _, jw := range sendJwItems {
		if service.CheckTime(jw.Date, config.Configs.Send.Range) {
			sendItems = append(sendItems, jw)
		}
	}
	for _, se := range sendSeItems {
		if service.CheckTime(se.Date, config.Configs.Send.Range) {
			sendItems = append(sendItems, se)
		}
	}
	if len(sendItems) == 0 {
		return
	}

	// send items
	moreStr := fmt.Sprintf(
		"--- \n+ 更多信息，请查阅 [华工教务通知](%s) 以及 [软院公务通知](%s)。",
		static.JwHomepage, static.SeHomepage,
	)

	// send bot
	sb := strings.Builder{}
	sb.WriteString("*教务系统通知*\n=====\n")
	for _, item := range sendItems {
		sb.WriteString(fmt.Sprintf("+ %s\n", item.String()))
	}
	sb.WriteString(moreStr)
	msg := sb.String()
	err = bot.SendToChat(user.ChatID, msg, telebot.ModeMarkdown)
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
			sb.WriteString(fmt.Sprintf("+ %s\n", sendItems[j].String()))
		}
		if i == sendTimes-1 {
			sb.WriteString(moreStr)
		}
		msg := sb.String()

		title := fmt.Sprintf("教务系统通知 (第 %d 条，共 %d 条)", i+1, sendTimes)
		_ = wechat.SendToChat(user.Sckey, title, msg)
	}

	// write old data
	oldJwData[user.ChatID] = newJwItems
	oldSeData[user.ChatID] = newSeItems
}
