package controller

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib-web/xtelebot"
	"github.com/Aoi-hosizora/ahlib/xstatus"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/service/dao"
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

// StartCtrl /start
func StartCtrl(bw *xtelebot.BotWrapper, m *telebot.Message) {
	bw.ReplyTo(m, START)
}

// HelpCtrl /help
func HelpCtrl(bw *xtelebot.BotWrapper, m *telebot.Message) {
	bw.ReplyTo(m, HELP, telebot.ModeMarkdown)
}

// SubscribeCtrl /subscribe
func SubscribeCtrl(bw *xtelebot.BotWrapper, m *telebot.Message) {
	sts := dao.CreateUser(m.Chat.ID)
	flag := ""
	if sts == xstatus.DbExisted {
		flag = SUBSCRIBE_ALREADY
	} else if sts == xstatus.DbFailed {
		flag = SUBSCRIBE_FAILED
	} else {
		flag = SUBSCRIBE_SUCCESS
	}
	bw.ReplyTo(m, flag)
}

// UnsubscribeCtrl /unsubscribe
func UnsubscribeCtrl(bw *xtelebot.BotWrapper, m *telebot.Message) {
	sts := dao.DeleteUser(m.Chat.ID)
	flag := ""
	if sts == xstatus.DbNotFound {
		flag = SUBSCRIBE_NOT_YET
	} else if sts == xstatus.DbFailed {
		flag = UNSUBSCRIBE_FAILED
	} else {
		flag = UNSUBSCRIBE_SUCCESS
	}
	bw.ReplyTo(m, flag)
}

// OnTextCtrl $on_text
func OnTextCtrl(bw *xtelebot.BotWrapper, m *telebot.Message) {
	bw.ReplyTo(m, fmt.Sprintf(UNKNOWN_COMMAND, m.Text))
}
