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
	SUBSCRIBE_ALREADY = "你已经绑定通知器至当前会话，请先发送 /unsubscribe 取消绑定。"
	SUBSCRIBE_FAILED  = "通知器绑定失败，请稍后重试。"
	SUBSCRIBE_SUCCESS = "通知器绑定成功，你将会在相关通知更新时收到推送。"

	SUBSCRIBE_NOT_YET   = "你还没有绑定通知器到当前会话，请先发送 /subscribe 绑定。"
	UNSUBSCRIBE_FAILED  = "通知器取消绑定失败，请稍后重试。"
	UNSUBSCRIBE_SUCCESS = "通知器取消绑定成功，你将不会收到通知推送。"

	GET_DATA_FAILED = "无法获取通知内容，请稍后重试。"
	NO_NEW_DATA     = "当前没有通知内容。"
)

// SubscribeCtrl /subscribe
func SubscribeCtrl(bw *xtelebot.BotWrapper, m *telebot.Message) {
	sts := dao.CreateChat(m.Chat.ID)
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
	sts := dao.DeleteChat(m.Chat.ID)
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

// SendCtrl /send
func SendCtrl(bw *xtelebot.BotWrapper, m *telebot.Message) {
	items, err := service.GetNoticeItems()
	if err != nil {
		if config.IsDebugMode() {
			bw.ReplyTo(m, GET_DATA_FAILED+"\n错误："+err.Error())
		} else {
			bw.ReplyTo(m, GET_DATA_FAILED)
		}
		return
	}
	if len(items) == 0 {
		bw.ReplyTo(m, NO_NEW_DATA)
		return
	}

	rendered := service.RenderNoticeItems(items, false)
	bw.ReplyTo(m, rendered, telebot.ModeMarkdown)
}
