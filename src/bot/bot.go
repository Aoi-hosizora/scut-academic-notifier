package bot

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/bot/controller"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/bot/server"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/config"
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

func Setup() error {
	b, err := telebot.NewBot(telebot.Settings{
		Token:   config.Configs.Bot.Token,
		Verbose: false,
		Poller: &telebot.LongPoller{
			Timeout: time.Second * time.Duration(config.Configs.Bot.PollerTimeout),
		},
	})
	if err != nil {
		return err
	}

	fmt.Println()
	log.Printf("[Telebot] Success to connect telegram bot: @%s\n", b.Me.Username)
	fmt.Println()

	server.Bot = server.NewBotServer(b)
	initHandler(server.Bot)

	return nil
}

func initHandler(b *server.BotServer) {
	// start
	b.HandleMessage("/start", controller.StartCtrl)
	b.HandleMessage("/help", controller.HelpCtrl)
	b.HandleMessage("/cancel", controller.CancelCtrl)
	b.HandleMessage(telebot.OnText, controller.OnTextCtrl)

	// wechat
	b.InlineButtons["btn_unbind"] = &telebot.InlineButton{Unique: "btn_unbind", Text: "Unbind"}
	b.InlineButtons["btn_cancel"] = &telebot.InlineButton{Unique: "btn_cancel", Text: "Cancel"}
	b.HandleMessage("/bind", controller.BindCtrl)
	b.HandleMessage("/unbind", controller.UnbindCtrl)
	b.HandleInline(b.InlineButtons["btn_unbind"], controller.InlUnbindBtnCtrl)
	b.HandleInline(b.InlineButtons["btn_cancel"], controller.InlCancelBtnCtrl)

	// notifier
	b.HandleMessage("/send", controller.SendCtrl)
}

func SendToChat(chatId int64, what interface{}, options ...interface{}) error {
	chat, err := server.Bot.Bot.ChatByID(xnumber.FormatInt64(chatId, 10))
	if err != nil {
		return err
	}

	return server.Bot.Send(chat, what, options...)
}
