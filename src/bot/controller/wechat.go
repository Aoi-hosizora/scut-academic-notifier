package controller

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xstatus"
	"github.com/Aoi-hosizora/go-serverchan"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/bot/button"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/bot/fsm"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/bot/server"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/database"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/model"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/wechat"
	"gopkg.in/tucnak/telebot.v2"
	"strings"
)

// /bind
func BindCtrl(m *telebot.Message) {
	user := database.GetUser(m.Chat.ID)
	if user != nil {
		_ = server.Bot.Reply(m, BIND_ALREADY)
	} else {
		server.Bot.SetStatus(m.Chat.ID, fsm.Binding)
		_ = server.Bot.Reply(m, BIND_Q)
	}
}

// /bind -> x
func fromBindingCtrl(m *telebot.Message) {
	sckey := strings.TrimSpace(m.Text)
	ok, err := wechat.CheckSckey(sckey, "A test message for binding by telebot")
	if err != nil {
		server.Bot.SetStatus(m.Chat.ID, fsm.None)
		_ = server.Bot.Reply(m, BIND_FAILED)
		return
	}
	if !ok {
		_ = server.Bot.Reply(m, SCKEY_INVALID)
		return
	}

	user := &model.User{ChatID: m.Chat.ID, Sckey: sckey}
	status := database.AddUser(user)

	server.Bot.SetStatus(m.Chat.ID, fsm.None)
	if status == xstatus.DbExisted {
		_ = server.Bot.Reply(m, BIND_ALREADY)
	} else if status == xstatus.DbFailed {
		_ = server.Bot.Reply(m, BIND_FAILED)
	} else {
		_ = server.Bot.Reply(m, BIND_SUCCESS)
	}
}

// /unbind
func UnbindCtrl(m *telebot.Message) {
	user := database.GetUser(m.Chat.ID)
	if user == nil {
		_ = server.Bot.Reply(m, BIND_NOT_YET)
	} else {
		_ = server.Bot.Reply(m, fmt.Sprintf(UNBIND_Q, serverchan.Mask(user.Sckey)), &telebot.ReplyMarkup{
			InlineKeyboard: [][]telebot.InlineButton{
				{*button.InlineBtnUnbind}, {*button.InlineBtnCancel},
			},
		})
	}
}

// inl:btn_unbind
func InlUnbindBtnCtrl(c *telebot.Callback) {
	m := c.Message
	_ = server.Bot.Delete(m)

	status := database.DeleteUser(m.Chat.ID)

	if status == xstatus.DbNotFound {
		_ = server.Bot.Reply(m, BIND_NOT_YET)
	} else if status == xstatus.DbFailed {
		_ = server.Bot.Reply(m, UNBIND_FAILED)
	} else {
		_ = server.Bot.Reply(m, UNBIND_SUCCESS)
	}
}

// inl:btn_cancel
func InlCancelBtnCtrl(c *telebot.Callback) {
	m := c.Message
	_ = server.Bot.Delete(m)

	_ = server.Bot.Reply(m, ACTION_CANCELED)
}
