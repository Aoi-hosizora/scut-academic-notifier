package controller

import (
	"fmt"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/bot/server"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/model"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/config"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/dao"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/service"
	"gopkg.in/tucnak/telebot.v2"
	"strings"
)

// /send
func SendCtrl(m *telebot.Message) {
	user := dao.GetUser(m.Chat.ID)
	if user == nil {
		_ = server.Bot().Reply(m, BIND_NOT_YET)
		return
	}

	jwItems, err := service.GetJwItems()
	if err != nil {
		_ = server.Bot().Reply(m, GET_DATA_FAILED)
		return
	}
	seItems, err := service.GetSeItems()
	if err != nil {
		_ = server.Bot().Reply(m, GET_DATA_FAILED)
		return
	}

	items := make([]*model.PostItem, 0)
	for _, jw := range jwItems {
		if service.CheckTime(jw.Date, config.Configs().Send.Range) {
			items = append(items, jw)
		}
	}
	for _, se := range seItems {
		if service.CheckTime(se.Date, config.Configs().Send.Range) {
			items = append(items, se)
		}
	}
	if len(items) == 0 {
		_ = server.Bot().Reply(m, NO_NEW_DATA)
		return
	}

	sb := strings.Builder{}
	sb.WriteString("*学校相关通知*\n=====\n")
	for idx, item := range items {
		sb.WriteString(fmt.Sprintf("%d. %s\n", idx+1, item.String()))
	}
	sb.WriteString(fmt.Sprintf(
		"===== \n更多信息，请查阅 [华工教务通知](%s) 以及 [软院公务通知](%s)。",
		service.JwHomepage, service.SeHomepage,
	))

	msg := sb.String()
	_ = server.Bot().Reply(m, msg, telebot.ModeMarkdown)
}
