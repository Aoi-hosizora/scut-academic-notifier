package main

import (
	"flag"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/config"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/server"
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
	cfg, err := config.LoadConfig(*fConfig)
	if err != nil {
		log.Fatalln("Failed to load config file:", err)
	}
	static := config.LoadStatic()

	s := server.NewServer(cfg, static)
	s.Serve()
}
