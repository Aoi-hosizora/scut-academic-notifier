package controller

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xslice"
	"github.com/Aoi-hosizora/ahlib/xstatus"
	"github.com/Aoi-hosizora/ahlib/xstring"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/bot/button"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/bot/fsm"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/bot/server"
	serverchan2 "github.com/Aoi-hosizora/scut-academic-notifier/internal/bot/serverchan"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/model"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/dao"
	"gopkg.in/tucnak/telebot.v2"
	"strings"
)

// /bind
func BindCtrl(m *telebot.Message) {
	user := dao.GetUser(m.Chat.ID)
	if user != nil {
		_ = server.Bot().Reply(m, BIND_ALREADY)
	} else {
		server.Bot().SetStatus(m.Chat.ID, fsm.Binding)
		_ = server.Bot().Reply(m, BIND_Q)
	}
}

// fsm.Binding
func FromBindingCtrl(m *telebot.Message) {
	sckey := strings.TrimSpace(m.Text)
	ok, err := serverchan2.CheckSckey(sckey, "A test message for binding by telebot")
	if err != nil {
		server.Bot().SetStatus(m.Chat.ID, fsm.None)
		_ = server.Bot().Reply(m, BIND_FAILED)
		return
	}
	if !ok {
		_ = server.Bot().Reply(m, SCKEY_INVALID)
		return
	}

	user := &model.User{ChatID: m.Chat.ID, Sckey: sckey}
	status := dao.AddUser(user)

	server.Bot().SetStatus(m.Chat.ID, fsm.None)
	if status == xstatus.DbExisted {
		_ = server.Bot().Reply(m, BIND_ALREADY)
	} else if status == xstatus.DbFailed {
		_ = server.Bot().Reply(m, BIND_FAILED)
	} else {
		_ = server.Bot().Reply(m, BIND_SUCCESS)
	}
}

// /unbind
func UnbindCtrl(m *telebot.Message) {
	user := dao.GetUser(m.Chat.ID)
	if user == nil {
		_ = server.Bot().Reply(m, BIND_NOT_YET)
	} else {
		indices := append(xslice.Range(6, 15, 1), append(xslice.Range(22, 31, 1), xslice.Range(38, 47, 1)...)...)
		maskedSckey := xstring.MaskToken(user.Sckey, '*', indices...)
		_ = server.Bot().Reply(m, fmt.Sprintf(UNBIND_Q, maskedSckey), &telebot.ReplyMarkup{
			InlineKeyboard: [][]telebot.InlineButton{
				{*button.InlineBtnUnbind}, {*button.InlineBtnCancel},
			},
		})
	}
}

// button.InlineBtnUnbind
func InlineBtnUnbindCtrl(c *telebot.Callback) {
	m := c.Message
	_, _ = server.Bot().Edit(m, fmt.Sprintf("%s (unbind)", m.Text))

	status := dao.DeleteUser(m.Chat.ID)

	if status == xstatus.DbNotFound {
		_ = server.Bot().Reply(m, BIND_NOT_YET)
	} else if status == xstatus.DbFailed {
		_ = server.Bot().Reply(m, UNBIND_FAILED)
	} else {
		_ = server.Bot().Reply(m, UNBIND_SUCCESS)
	}
}

// button.InlineBtnCancel
func InlineBtnCancelCtrl(c *telebot.Callback) {
	m := c.Message
	_, _ = server.Bot().Edit(m, fmt.Sprintf("%s (canceled)", m.Text))
}
