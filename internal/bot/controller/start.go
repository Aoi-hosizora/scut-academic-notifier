package controller

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib-web/xtelebot"
	"gopkg.in/tucnak/telebot.v2"
)

const (
	START = "这里是 @AoiHosizora 开发的华工教务通知器，通知内容包括：华工教务处通知、华工软院公务通知，发送 /help 显示帮助信息。"

	HELP = `*开始命令*
/start - 显示开始信息
/help - 显示帮助信息

*通知器命令*
/subscribe - 订阅通知器
/unsubscribe - 取消订阅通知器
/send - 发送最新通知列表

*Bug 反馈*
https://github.com/Aoi-hosizora/scut-academic-notifier/issues`

	UNKNOWN_COMMAND = "未知命令：%s，发送 /help 显示帮助信息。"
)

// StartCtrl /start
func StartCtrl(bw *xtelebot.BotWrapper, m *telebot.Message) {
	bw.ReplyTo(m, START)
}

// HelpCtrl /help
func HelpCtrl(bw *xtelebot.BotWrapper, m *telebot.Message) {
	bw.ReplyTo(m, HELP, telebot.ModeMarkdown)
}

// OnTextCtrl $on_text
func OnTextCtrl(bw *xtelebot.BotWrapper, m *telebot.Message) {
	bw.ReplyTo(m, fmt.Sprintf(UNKNOWN_COMMAND, m.Text))
}
