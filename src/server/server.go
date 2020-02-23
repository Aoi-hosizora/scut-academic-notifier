package server

import (
	"fmt"
	"github.com/Aoi-hosizora/SCUT_Academic_Notifier/src/config"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Server struct {
	Config *config.Config
	Static *config.Static
}

func NewServer(config *config.Config, static *config.Static) *Server {
	return &Server{
		Config: config,
		Static: static,
	}
}

func (s *Server) send(title string, message string) bool {
	title = url.QueryEscape(title)
	message = url.QueryEscape(message)
	sendUrl := fmt.Sprintf(s.Static.ServerChanUrl, s.Config.WxConfig.Sckey, title, message)

	resp, err := http.Post(sendUrl, "application/x-www-form-urlencoded", strings.NewReader("name=cjb"))
	if err != nil {
		log.Println("Failed to post data:", err)
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("Success to send but get %d response", resp.StatusCode)
		return false
	}
	return true
}

func (s *Server) Serve() {
	if s.Config.WxConfig.Sckey == "" {
		log.Fatalln("Sckey could not be empty")
	}
	log.Printf("Start listening, pollng every %d second(s)...\n", s.Config.ServerConfig.PollingDuration)
	s.polling()
}
