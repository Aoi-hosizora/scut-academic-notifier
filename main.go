package main

import (
	"flag"
	"fmt"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/bot"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/bot/server"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/config"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/database"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/logger"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/task"
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
	err = database.SetupGorm()
	if err != nil {
		log.Fatalln("Failed to connect mysql:", err)
	}
	err = database.SetupRedis()
	if err != nil {
		log.Fatalln("Failed to connect redis:", err)
	}
	fmt.Println()
	err = bot.Setup()
	if err != nil {
		log.Fatalln("Failed to load telebot:", err)
	}
	err = wechat.Setup()
	if err != nil {
		log.Fatalln("Failed to load serverchan:", err)
	}
	err = task.Setup()
	if err != nil {
		log.Fatalln("Failed to load cron:", err)
	}

	defer task.Cron.Stop()
	task.Cron.Start()

	defer server.Bot.Stop()
	server.Bot.Start()
}
