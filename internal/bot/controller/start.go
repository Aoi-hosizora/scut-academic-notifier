package controller

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib-web/xtelebot"
	"gopkg.in/tucnak/telebot.v2"
)

const (
	_START = "这里是 @AoiHosizora 开发的华工教务通知器，通知内容来自：教务处、软院、研究生院、GZIC，发送 /help 显示帮助信息。"

	_HELP = `*开始命令*
/start - 显示开始信息
/help - 显示帮助信息

*通知器命令*
/subscribe - 订阅通知器
/unsubscribe - 取消订阅通知器
/links - 获取通知来源链接
/send - 获取最新通知内容

*Bug 反馈*
https://github.com/Aoi-hosizora/scut-academic-notifier/issues`

	_UNKNOWN_COMMAND = "未知命令：%s，发送 /help 显示帮助信息。"
)

// StartCtrl /start
func StartCtrl(bw *xtelebot.BotWrapper, m *telebot.Message) {
	bw.ReplyTo(m, _START)
}

// HelpCtrl /help
func HelpCtrl(bw *xtelebot.BotWrapper, m *telebot.Message) {
	bw.ReplyTo(m, _HELP, telebot.ModeMarkdown)
}

// OnTextCtrl $on_text
func OnTextCtrl(bw *xtelebot.BotWrapper, m *telebot.Message) {
	bw.ReplyTo(m, fmt.Sprintf(_UNKNOWN_COMMAND, m.Text))
}
