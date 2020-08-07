package controller

import (
	"github.com/Aoi-hosizora/scut-academic-notifier/src/bot/fsm"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/bot/server"
	"gopkg.in/tucnak/telebot.v2"
)

// /start
func StartCtrl(m *telebot.Message) {
	_ = server.Bot.Reply(m, "Here is aoihosizora's scut academic notifier bot. See /help for help.")
}

// /help
func HelpCtrl(m *telebot.Message) {
	_ = server.Bot.Reply(m, `*Commands*
/start - show start message
/help - show this help message
/cancel - cancel the last action`, telebot.ModeMarkdown)
}

// /cancel
func CancelCtrl(m *telebot.Message) {
	if server.Bot.UsersData.GetStatus(m.Chat.ID) == fsm.None {
		_ = server.Bot.Reply(m, "There is no action now.")
	} else {
		server.Bot.UsersData.SetStatus(m.Chat.ID, fsm.None)
		_ = server.Bot.Reply(m, "Action has been canceled.", &telebot.ReplyMarkup{
			ReplyKeyboardRemove: true,
		})
	}
}

// $onText
func OnTextCtrl(m *telebot.Message) {
	switch server.Bot.UsersData.GetStatus(m.Chat.ID) {
	default:
		_ = server.Bot.Reply(m, "Unknown command: "+m.Text+".")
	}
}
