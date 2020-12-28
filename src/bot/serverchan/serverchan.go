package serverchan

import (
	"github.com/Aoi-hosizora/go-serverchan"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/logger"
)

var Serverchan *serverchan.Client

func Setup() error {
	Serverchan = serverchan.NewClient()
	Serverchan.SetLogger(logger.Serverchan)
	return nil
}

func SendToChat(sckey string, title string, message string) error {
	_, _, err := Serverchan.Send(sckey, title, message)
	return err
}

func CheckSckey(sckey string, title string) (bool, error) {
	return Serverchan.CheckSckey(sckey, title)
}
