package main

import (
	"flag"
	"fmt"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/bot"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/bot/server"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/config"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/logger"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/wechat"
	"log"
)

var (
	hHelp   = flag.Bool("h", false, "show help")
	fConfig = flag.String("config", "./config.yaml", "config path")
)

func main() {
	flag.Parse()
	if *hHelp {
		flag.Usage()
	} else {
		run()
	}
}

func run() {
	err := config.Load(*fConfig)
	if err != nil {
		log.Fatalln("Failed to load config:", err)
	}
	err = logger.Setup()
	if err != nil {
		log.Fatalln("Failed to setup logger:", err)
	}
	fmt.Println()
	err = bot.Setup()
	if err != nil {
		log.Fatalln("Failed to load telebot:", err)
	}
	err = wechat.Setup()
	if err != nil {
		log.Fatalln("Failed to load wechat server-chan:", err)
	}

	defer server.Bot.Stop()
	server.Bot.Start()
}
