package serverchan

import (
	"github.com/Aoi-hosizora/go-serverchan"
)

var Serverchan *serverchan.Client

func Setup() error {
	Serverchan = serverchan.NewClient()
	return nil
}

func SendToChat(sckey string, title string, message string) error {
	return Serverchan.Send(sckey, title, message)
}

func CheckSckey(sckey string, title string) (bool, error) {
	return Serverchan.CheckSckey(sckey, title)
}
