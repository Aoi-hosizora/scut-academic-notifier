package controller

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xstatus"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/bot/server"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/dao"
	"gopkg.in/tucnak/telebot.v2"
)

const (
	START = "Here is AoiHosizora's scut academic notifier bot, developed by @AoiHosizora. See /help to show help message."
	HELP  = `*Commands*
/start - show start message
/help - show this help message
/subscribe - subscribe this chat
/unsubscribe - unsubscribe this chat

*Notifier*
/send - send the current notices

*Bug report*
https://github.com/Aoi-hosizora/scut-academic-notifier/issues/new`

	UNKNOWN_COMMAND = "Unknown command: %s. See /help for help."

	SUBSCRIBE_ALREADY = "You have already subscribed this chat, use /unsubscribe to unsubscribe first."
	SUBSCRIBE_FAILED  = "Failed to subscribe, please retry later."
	SUBSCRIBE_SUCCESS = "Subscribe success, you will receive notifiers when updated."

	SUBSCRIBE_NOT_YET   = "You have not subscribe yet, use /subscribe to subscribe first."
	UNSUBSCRIBE_FAILED  = "Failed to unsubscribe, please retry later."
	UNSUBSCRIBE_SUCCESS = "Unsubscribe success, you will not receive any notifiers now."
)

// /start
func StartCtrl(m *telebot.Message) {
	_ = server.Bot().Reply(m, START)
}

// /help
func HelpCtrl(m *telebot.Message) {
	_ = server.Bot().Reply(m, HELP, telebot.ModeMarkdown)
}

// /subscribe
func SubscribeCtrl(m *telebot.Message) {
	sts := dao.CreateUser(m.Chat.ID)
	flag := ""
	if sts == xstatus.DbExisted {
		flag = SUBSCRIBE_ALREADY
	} else if sts == xstatus.DbFailed {
		flag = SUBSCRIBE_FAILED
	} else {
		flag = SUBSCRIBE_SUCCESS
	}
	_ = server.Bot().Reply(m, flag)
}

// /unsubscribe
func UnsubscribeCtrl(m *telebot.Message) {
	sts := dao.DeleteUser(m.Chat.ID)
	flag := ""
	if sts == xstatus.DbNotFound {
		flag = SUBSCRIBE_NOT_YET
	} else if sts == xstatus.DbFailed {
		flag = UNSUBSCRIBE_FAILED
	} else {
		flag = UNSUBSCRIBE_SUCCESS
	}
	_ = server.Bot().Reply(m, flag)
}

// $on_text
func OnTextCtrl(m *telebot.Message) {
	_ = server.Bot().Reply(m, fmt.Sprintf(UNKNOWN_COMMAND, m.Text))
}
