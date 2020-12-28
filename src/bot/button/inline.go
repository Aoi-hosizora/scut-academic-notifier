package button

import (
	"gopkg.in/tucnak/telebot.v2"
)

var (
	// Used for controller.UnbindCtrl and controller.AllowIssueCtrl.
	// Callback: controller.InlBtnCancelCtrl.
	InlineBtnCancel = &telebot.InlineButton{Unique: "btn_cancel", Text: "Cancel"}

	// Used for controller.UnbindCtrl.
	// Callback: controller.InlBtnUnbindCtrl.
	InlineBtnUnbind = &telebot.InlineButton{Unique: "btn_unbind", Text: "Unbind"}
)
