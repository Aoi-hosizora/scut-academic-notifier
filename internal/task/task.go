package task

import (
	"fmt"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/bot/server"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/model"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/config"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/dao"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/logger"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/service"
	"github.com/robfig/cron/v3"
	"gopkg.in/tucnak/telebot.v2"
	"strings"
	"sync"
)

// _cron represents the global cron.Cron.
var _cron *cron.Cron

func Cron() *cron.Cron {
	return _cron
}

func Setup() error {
	cr := cron.New(cron.WithSeconds())

	_, err := cr.AddFunc(config.Configs().Task.Cron, task)
	if err != nil {
		return err
	}

	_cron = cr
	return nil
}

func foreachUsers(users []*model.User, fn func(user *model.User)) {
	wg := sync.WaitGroup{}
	for _, user := range users {
		wg.Add(1)
		go func(user *model.User) {
			defer func() { recover() }()
			fn(user)
			wg.Done()
		}(user)
	}
	wg.Wait()
}

func task() {
	defer func() { recover() }()

	users := dao.QueryUsers()
	if len(users) == 0 {
		return
	}

	foreachUsers(users, func(user *model.User) {
		// get new items
		jwItems, err := service.GetJwItems()
		if err != nil || len(jwItems) == 0 {
			return
		}
		seItems, err := service.GetSeItems()
		if err != nil || len(seItems) == 0 {
			return
		}

		// filter new items
		newItems := make([]*model.PostItem, 0)
		for _, jw := range jwItems {
			if service.CheckInTimeRange(jw.Date, config.Configs().Send.TimeRange) {
				newItems = append(newItems, jw)
			}
		}
		for _, se := range seItems {
			if service.CheckInTimeRange(se.Date, config.Configs().Send.TimeRange) {
				newItems = append(newItems, se)
			}
		}
		if len(newItems) == 0 {
			return
		}

		// get old items and calc diff
		oldItems, ok := dao.GetOldItems(user.ChatID)
		if !ok {
			return
		}
		logger.Logger().Infof("Get old items: #%d | %d", len(oldItems), user.ChatID)
		diff := model.ItemSliceDiff(newItems, oldItems)
		logger.Logger().Infof("Get diff items: #%d | %d", len(diff), user.ChatID)
		if len(diff) == 0 {
			return
		}

		// update old items
		ok = dao.SetOldItems(user.ChatID, newItems)
		if !ok {
			return
		}
		logger.Logger().Infof("Set new items: #%d | %d", len(newItems), user.ChatID)

		// render
		sb := strings.Builder{}
		sb.WriteString("*学校相关通知*\n=====\n")
		for idx, item := range diff {
			sb.WriteString(fmt.Sprintf("%d. %s\n", idx+1, item.String()))
		}
		footer := fmt.Sprintf("=====\n更多信息，请查阅 [华工教务通知](%s) 以及 [软院公务通知](%s)。", service.JwHomepage, service.SeHomepage)
		sb.WriteString(footer)
		render := sb.String()

		// send
		_ = server.Bot().SendToChat(user.ChatID, render, telebot.ModeMarkdown)
	})
}
