package bot

import (
	"github.com/Aoi-hosizora/ahlib-web/xrecovery"
	"github.com/Aoi-hosizora/ahlib/xruntime"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/bot/button"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/bot/controller"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/bot/server"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/config"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/logger"
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

func Setup() error {
	b, err := telebot.NewBot(telebot.Settings{
		Token:   config.Configs().Bot.Token,
		Verbose: false,
		Poller: &telebot.LongPoller{
			Timeout: time.Second * time.Duration(config.Configs().Bot.PollerTimeout),
		},
		Reporter: func(err error) {
			xrecovery.LogToLogrus(logger.Logger(), err, xruntime.RuntimeTraceStack(0))
		},
	})
	if err != nil {
		return err
	}

	log.Println("Success to connect telegram bot:", b.Me.Username)
	server.SetupBot(server.NewBotServer(b))
	setupHandler(server.Bot())

	return nil
}

func setupHandler(b *server.BotServer) {
	// start
	b.HandleMessage("/start", controller.StartCtrl)
	b.HandleMessage("/help", controller.HelpCtrl)
	b.HandleMessage("/cancel", controller.CancelCtrl)
	b.HandleMessage(telebot.OnText, controller.OnTextCtrl)

	// wechat
	b.HandleMessage("/bind", controller.BindCtrl)
	b.HandleMessage("/unbind", controller.UnbindCtrl)
	b.HandleInline(button.InlineBtnUnbind, controller.InlineBtnUnbindCtrl)
	b.HandleInline(button.InlineBtnCancel, controller.InlineBtnCancelCtrl)

	// notifier
	b.HandleMessage("/send", controller.SendCtrl)
}
