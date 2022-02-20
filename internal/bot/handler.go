package bot

import (
	"github.com/Aoi-hosizora/ahlib-web/xtelebot"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/bot/controller"
	"gopkg.in/tucnak/telebot.v2"
)

func setupHandlers(bw *xtelebot.BotWrapper) {
	// start
	bw.HandleCommand("/start", controller.Start)
	bw.HandleCommand("/help", controller.Help(help))
	bw.HandleCommand(telebot.OnText, controller.OnText)

	// notifier
	bw.HandleCommand("/subscribe", controller.Subscribe)
	bw.HandleCommand("/unsubscribe", controller.Unsubscribe)
	bw.HandleCommand("/links", controller.Links)
	bw.HandleCommand("/send", controller.Send)

	// set commands
	_ = bw.Bot().SetCommands(commands)
}

const help = `*开始命令*
/start - 显示开始信息
/help - 显示帮助信息

*通知器命令*
/subscribe - 订阅通知器
/unsubscribe - 取消订阅通知器
/links - 获取通知来源链接
/send - 获取最新通知内容

*Bug 反馈*
https://github.com/Aoi-hosizora/scut-academic-notifier/issues`

var commands = []telebot.Command{
	{"/start", "显示开始信息"},
	{"/help", "显示帮助信息"},
	{"/subscribe", "订阅通知器"},
	{"/unsubscribe", "取消订阅通知器"},
	{"/links", "获取通知来源链接"},
	{"/send", "获取最新通知内容"},
}
