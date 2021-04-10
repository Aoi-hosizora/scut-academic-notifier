package button

import (
	"gopkg.in/tucnak/telebot.v2"
)

var (
	// controller.UnbindCtrl -> controller.InlineBtnCancelCtrl
	InlineBtnCancel = &telebot.InlineButton{Unique: "btn_cancel", Text: "Cancel"}

	// controller.UnbindCtrl -> controller.InlineBtnUnbindCtrl
	InlineBtnUnbind = &telebot.InlineButton{Unique: "btn_unbind", Text: "Unbind"}
)
