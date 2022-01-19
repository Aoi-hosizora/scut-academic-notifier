package controller

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib-web/xtelebot"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/model"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/config"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/service"
	"gopkg.in/tucnak/telebot.v2"
	"strings"
)

const (
	GET_DATA_FAILED = "Failed to get notice information, please retry later."
	NO_NEW_DATA     = "There is no notice."
)

// SendCtrl /send
func SendCtrl(bw *xtelebot.BotWrapper, m *telebot.Message) {
	jwItems, err := service.GetJwItems()
	if err != nil {
		bw.ReplyTo(m, GET_DATA_FAILED)
		return
	}
	seItems, err := service.GetSeItems()
	if err != nil {
		bw.ReplyTo(m, GET_DATA_FAILED)
		return
	}

	items := make([]*model.PostItem, 0)
	for _, jw := range jwItems {
		if service.CheckInTimeRange(jw.Date, config.Configs().Task.NotifierTimeRange) {
			items = append(items, jw)
		}
	}
	for _, se := range seItems {
		if service.CheckInTimeRange(se.Date, config.Configs().Task.NotifierTimeRange) {
			items = append(items, se)
		}
	}
	if len(items) == 0 {
		bw.ReplyTo(m, NO_NEW_DATA)
		return
	}

	sb := strings.Builder{}
	sb.WriteString("*学校相关通知*\n=====\n")
	for idx, item := range items {
		sb.WriteString(fmt.Sprintf("%d. %s\n", idx+1, item.String()))
	}
	footer := fmt.Sprintf("=====\n更多信息，请查阅 [华工教务通知](%s) 以及 [软院公务通知](%s)。", service.JwHomepage, service.SeHomepage)
	sb.WriteString(footer)

	msg := sb.String()
	bw.ReplyTo(m, msg, telebot.ModeMarkdown)
}
