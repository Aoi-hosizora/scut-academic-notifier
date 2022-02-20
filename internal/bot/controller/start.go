package controller

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib-web/xtelebot"
	"gopkg.in/tucnak/telebot.v2"
)

const (
	_START           = "这里是 @AoiHosizora 开发的华工教务通知器，通知内容来自：教务处、软件学院、研究生院、GZIC，发送 /help 显示帮助信息。"
	_UNKNOWN_COMMAND = "未知命令：%s，发送 /help 显示帮助信息。"
)

// Start /start
func Start(bw *xtelebot.BotWrapper, m *telebot.Message) {
	bw.RespondReply(m, false, _START)
}

// Help /help
func Help(help string) xtelebot.MessageHandler {
	return func(bw *xtelebot.BotWrapper, m *telebot.Message) {
		bw.RespondReply(m, false, help, telebot.ModeMarkdown)
	}
}

// OnText $on_text
func OnText(bw *xtelebot.BotWrapper, m *telebot.Message) {
	bw.RespondReply(m, false, fmt.Sprintf(_UNKNOWN_COMMAND, m.Text))
}
