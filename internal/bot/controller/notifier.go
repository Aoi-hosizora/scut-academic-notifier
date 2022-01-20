package controller

import (
	"github.com/Aoi-hosizora/ahlib-web/xtelebot"
	"github.com/Aoi-hosizora/ahlib/xstatus"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/config"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/service"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/service/dao"
	"gopkg.in/tucnak/telebot.v2"
)

const (
	_SUBSCRIBE_ALREADY = "你已经绑定通知器至当前会话，请先发送 /unsubscribe 取消绑定。"
	_SUBSCRIBE_FAILED  = "通知器绑定失败，请稍后重试。"
	_SUBSCRIBE_SUCCESS = "通知器绑定成功，你将会在相关通知更新时收到推送。"

	_SUBSCRIBE_NOT_YET   = "你还没有绑定通知器到当前会话，请先发送 /subscribe 绑定。"
	_UNSUBSCRIBE_FAILED  = "通知器取消绑定失败，请稍后重试。"
	_UNSUBSCRIBE_SUCCESS = "通知器取消绑定成功，你将不会收到通知推送。"

	_GET_NOTICES_FAILED   = "无法获取通知内容，请稍后重试。"
	_NO_NOTICES_CURRENTLY = "当前没有通知内容。"
)

// SubscribeCtrl /subscribe
func SubscribeCtrl(bw *xtelebot.BotWrapper, m *telebot.Message) {
	sts, err := dao.CreateChat(m.Chat.ID)
	flag := ""
	if sts == xstatus.DbExisted {
		flag = _SUBSCRIBE_ALREADY
	} else if sts == xstatus.DbFailed {
		flag = _SUBSCRIBE_FAILED
		if config.IsDebugMode() {
			flag += "\n错误：" + err.Error()
		}
	} else {
		flag = _SUBSCRIBE_SUCCESS
	}
	bw.ReplyTo(m, flag)
}

// UnsubscribeCtrl /unsubscribe
func UnsubscribeCtrl(bw *xtelebot.BotWrapper, m *telebot.Message) {
	sts, err := dao.DeleteChat(m.Chat.ID)
	flag := ""
	if sts == xstatus.DbNotFound {
		flag = _SUBSCRIBE_NOT_YET
	} else if sts == xstatus.DbFailed {
		flag = _UNSUBSCRIBE_FAILED
		if config.IsDebugMode() {
			flag += "\n错误：" + err.Error()
		}
	} else {
		flag = _UNSUBSCRIBE_SUCCESS
	}
	bw.ReplyTo(m, flag)
}

// LinksCtrl /links
func LinksCtrl(bw *xtelebot.BotWrapper, m *telebot.Message) {
	flag := "*通知来源链接*\n" + service.GetNoticeLinks()
	bw.ReplyTo(m, flag, telebot.ModeMarkdown)
}

// SendCtrl /send
func SendCtrl(bw *xtelebot.BotWrapper, m *telebot.Message) {
	items, err := service.GetNoticeItems()
	if err != nil {
		if config.IsDebugMode() {
			bw.ReplyTo(m, _GET_NOTICES_FAILED+"\n错误："+err.Error())
		} else {
			bw.ReplyTo(m, _GET_NOTICES_FAILED)
		}
		return
	}
	if len(items) == 0 {
		bw.ReplyTo(m, _NO_NOTICES_CURRENTLY)
		return
	}

	rendered := service.RenderNoticeItems(items, false)
	bw.ReplyTo(m, rendered, telebot.ModeMarkdown)
}
