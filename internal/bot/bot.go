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

type Server struct {
	bot *xtelebot.BotWrapper
}

func (s *Server) Bot() *xtelebot.BotWrapper {
	return s.bot
}

func NewServer() (*Server, error) {
	// telebot
	bot, err := telebot.NewBot(telebot.Settings{
		Token:   config.Configs().Meta.Token,
		Verbose: false,
		Poller:  &telebot.LongPoller{Timeout: time.Second * time.Duration(config.Configs().Meta.PollerTimeout)},
	})
	if err != nil {
		return nil, err
	}
	log.Println("[Bot] Success to connect telegram bot:", bot.Me.Username)

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

	s := &Server{bot: bw}
	return s, nil
}

func setupLoggers(bw *xtelebot.BotWrapper) {
	l := logger.Logger()
	bw.SetReceivedCallback(func(endpoint interface{}, received *telebot.Message) {
		xtelebot.LogReceiveToLogrus(l, endpoint, received)
	})
	bw.SetAfterRepliedCallback(func(received *telebot.Message, replied *telebot.Message, err error) {
		xtelebot.LogReplyToLogrus(l, received, replied, err)
	})
	bw.SetAfterSentCallback(func(chat *telebot.Chat, sent *telebot.Message, err error) {
		xtelebot.LogSendToLogrus(l, chat, sent, err)
	})
	bw.SetPanicHandler(func(endpoint interface{}, v interface{}) {
		xgin.LogRecoveryToLogrus(l, v, xruntime.RuntimeTraceStack(4))
	})
}

func (s *Server) RunBot() {
	log.Println("[Bot] Starting consuming incoming bot update")
	s.bot.Bot().Start()
}

func setupHandlers(bw *xtelebot.BotWrapper) {
	// start
	bw.HandleCommand("/start", controller.StartCtrl)
	bw.HandleCommand("/help", controller.HelpCtrl)
	bw.HandleCommand("/subscribe", controller.SubscribeCtrl)
	bw.HandleCommand("/unsubscribe", controller.UnsubscribeCtrl)

	// notifier
	bw.HandleCommand("/send", controller.SendCtrl)

	// text
	bw.HandleCommand(telebot.OnText, controller.OnTextCtrl)
}
