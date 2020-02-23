package main

import (
	"flag"
	"github.com/Aoi-hosizora/SCUT_Academic_Notifier/src/config"
	"github.com/Aoi-hosizora/SCUT_Academic_Notifier/src/server"
	"log"
)

var (
	help bool
	configPath string
)

func init() {
	flag.BoolVar(&help, "h", false, "show help")
	flag.StringVar(&configPath, "config", "./src/config/Config.yaml", "config path")
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
	} else {
		run()
	}
}

func run() {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalln("Failed to load config file:", err)
	}
	static := config.LoadStatic()

	s := server.NewServer(cfg, static)
	s.Serve()
}
