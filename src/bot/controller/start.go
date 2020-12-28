package controller

import (
	"fmt"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/bot/fsm"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/bot/server"
	"gopkg.in/tucnak/telebot.v2"
)

// noinspection GoSnakeCaseUsage
const (
	START = "Here is aoihosizora's scut academic notifier bot. See /help for help."
	HELP  = `*Common*
/start - show start message
/help - show this help message
/cancel - cancel the last action

*Account*
/bind - bind with sckey
/unbind - unbind this chat

*Notifier*
/send - send the current notices

*Bug report*
https://github.com/Aoi-hosizora/scut-academic-notifier/issues/new
`

	NO_ACTION       = "There is no action now."
	ACTION_CANCELED = "Action has been canceled."
	UNKNOWN_COMMAND = "Unknown command: %s. See /help for help."

	BIND_ALREADY   = "You have already bound with a sckey."
	BIND_NOT_YET   = "You have not bound yet, use /bind to bind."
	BIND_Q         = "Please send a wechat sckey to bind. /cancel to cancel."
	SCKEY_INVALID  = "Sckey is invalid, please check and resend one."
	BIND_SUCCESS   = "Bind success, you will receive notifiers when updated."
	BIND_FAILED    = "Failed to bind, please retry later."
	UNBIND_Q       = "You have already bound with %s. Sure to unbind?"
	UNBIND_SUCCESS = "Unbind success, you will not receive any notifiers now."
	UNBIND_FAILED  = "Failed to unbind, please retry later."

	GET_DATA_FAILED = "Failed to get notice information, please retry later."
	NO_NEW_DATA     = "There is no new notice."
)

// /start
func StartCtrl(m *telebot.Message) {
	_ = server.Bot.Reply(m, START)
}

// /help
func HelpCtrl(m *telebot.Message) {
	_ = server.Bot.Reply(m, HELP, telebot.ModeMarkdown)
}

// /cancel
func CancelCtrl(m *telebot.Message) {
	if server.Bot.GetStatus(m.Chat.ID) == fsm.None {
		_ = server.Bot.Reply(m, NO_ACTION)
	} else {
		server.Bot.SetStatus(m.Chat.ID, fsm.None)
		_ = server.Bot.Reply(m, ACTION_CANCELED, &telebot.ReplyMarkup{
			ReplyKeyboardRemove: true,
		})
	}
}

// $onText
func OnTextCtrl(m *telebot.Message) {
	switch server.Bot.GetStatus(m.Chat.ID) {
	case fsm.Binding:
		fromBindingCtrl(m)
	default:
		_ = server.Bot.Reply(m, fmt.Sprintf(UNKNOWN_COMMAND, m.Text))
	}
}
