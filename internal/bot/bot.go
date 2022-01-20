package bot

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib-web/xgin"
	"github.com/Aoi-hosizora/ahlib-web/xtelebot"
	"github.com/Aoi-hosizora/ahlib/xcolor"
	"github.com/Aoi-hosizora/ahlib/xruntime"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/bot/controller"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/config"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/logger"
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

type Consumer struct {
	bw *xtelebot.BotWrapper
}

func (s *Consumer) BotWrapper() *xtelebot.BotWrapper {
	return s.bw
}

func NewConsumer() (*Consumer, error) {
	// telebot
	bot, err := telebot.NewBot(telebot.Settings{
		Token:   config.Configs().Meta.Token,
		Verbose: false,
		Poller:  &telebot.LongPoller{Timeout: time.Second * time.Duration(config.Configs().Meta.PollerTimeout)},
	})
	if err != nil {
		return nil, err
	}

	// wrapper
	bw := xtelebot.NewBotWrapper(bot)
	bw.SetEndpointHandledCallback(func(endpoint string, handlerName string) {
		if config.IsDebugMode() {
			fmt.Printf("[Bot-debug] %-46s --> %s\n", xcolor.Blue.Sprint(endpoint), handlerName)
		}
	})
	setupLoggers(bw)

	// handlers
	setupHandlers(bw)

	s := &Consumer{bw: bw}
	return s, nil
}

func setupLoggers(bw *xtelebot.BotWrapper) {
	l := logger.Logger()
	bw.SetReceivedCallback(func(endpoint interface{}, received *telebot.Message) {
		xtelebot.LogReceiveToLogrus(l, endpoint, received)
	})
	bw.SetRepliedCallback(func(received *telebot.Message, replied *telebot.Message, err error) {
		xtelebot.LogReplyToLogrus(l, received, replied, err)
	})
	bw.SetSentCallback(func(chat *telebot.Chat, sent *telebot.Message, err error) {
		xtelebot.LogSendToLogrus(l, chat, sent, err)
	})
	bw.SetPanicHandler(func(endpoint interface{}, v interface{}) {
		xgin.LogRecoveryToLogrus(l, v, xruntime.RuntimeTraceStack(4))
	})
}

func (s *Consumer) Start() {
	log.Printf("[Bot] Starting consuming incoming update on bot %s", s.bw.Bot().Me.Username)
	s.bw.Bot().Start() // block to poll and consume
}

func setupHandlers(bw *xtelebot.BotWrapper) {
	// start
	bw.HandleCommand("/start", controller.StartCtrl)
	bw.HandleCommand("/help", controller.HelpCtrl)
	bw.HandleCommand(telebot.OnText, controller.OnTextCtrl)

	// notifier
	bw.HandleCommand("/subscribe", controller.SubscribeCtrl)
	bw.HandleCommand("/unsubscribe", controller.UnsubscribeCtrl)
	bw.HandleCommand("/links", controller.LinksCtrl)
	bw.HandleCommand("/send", controller.SendCtrl)
}
