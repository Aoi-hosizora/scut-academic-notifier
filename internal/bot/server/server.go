package server

import (
	"github.com/Aoi-hosizora/ahlib-web/xtelebot"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/config"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/logger"
	"gopkg.in/tucnak/telebot.v2"
)

// _bot represents the global BotServer.
var _bot *BotServer

func Bot() *BotServer {
	return _bot
}

func SetupBot(b *BotServer) {
	_bot = b
}

type BotServer struct {
	bot  *telebot.Bot
	data *xtelebot.BotData
}

func NewBotServer(bot *telebot.Bot) *BotServer {
	return &BotServer{bot: bot, data: xtelebot.NewBotData()}
}

// ===
// Bot
// ===

func (b *BotServer) Start() {
	b.bot.Start()
}

func (b *BotServer) Stop() {
	b.bot.Stop()
}

func (b *BotServer) Edit(msg telebot.Editable, what interface{}, options ...interface{}) (*telebot.Message, error) {
	return b.bot.Edit(msg, what, options...)
}

func (b *BotServer) sendTo(c *telebot.Chat, what interface{}, cb func(*telebot.Message, error), options ...interface{}) error {
	var msg *telebot.Message
	var err error
	for i := 0; i < int(config.Configs().Bot.RetryCount); i++ { // retry
		msg, err = b.bot.Send(c, what, options...)
		cb(msg, err)
		if err == nil {
			return nil
		}
	}
	return err
}

func (b *BotServer) Send(c *telebot.Chat, what interface{}, options ...interface{}) error {
	return b.sendTo(c, what, func(msg *telebot.Message, err error) {
		logger.Send(c, msg, err)
	}, options...)
}

func (b *BotServer) Reply(m *telebot.Message, what interface{}, options ...interface{}) error {
	return b.sendTo(m.Chat, what, func(msg *telebot.Message, err error) {
		logger.Reply(m, msg, err)
	}, options...)
}

func (b *BotServer) SendToChat(chatId int64, what interface{}, options ...interface{}) error {
	chat, err := b.bot.ChatByID(xnumber.I64toa(chatId))
	if err != nil {
		return err
	}

	return b.Send(chat, what, options...)
}

// ======
// Handle
// ======

func (b *BotServer) HandleMessage(endpoint string, handler func(*telebot.Message)) {
	if handler == nil {
		panic("nil handler")
	}
	b.bot.Handle(endpoint, func(m *telebot.Message) {
		logger.Receive(endpoint, m)
		handler(m)
	})
}

func (b *BotServer) HandleInline(endpoint *telebot.InlineButton, handler func(*telebot.Callback)) {
	if handler == nil {
		panic("nil handler")
	}
	b.bot.Handle(endpoint, func(c *telebot.Callback) {
		logger.Receive(endpoint, c.Message)
		handler(c)
	})
}

func (b *BotServer) HandleReply(endpoint *telebot.ReplyButton, handler func(*telebot.Message)) {
	if handler == nil {
		panic("nil handler")
	}
	b.bot.Handle(endpoint, func(m *telebot.Message) {
		logger.Receive(endpoint, m)
		handler(m)
	})
}
