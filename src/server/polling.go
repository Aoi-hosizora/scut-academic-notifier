package server

import (
	"fmt"
	"github.com/Aoi-hosizora/SCUT_Academic_Notifier/src/model"
	"github.com/Aoi-hosizora/SCUT_Academic_Notifier/src/service"
	"github.com/Aoi-hosizora/ahlib/xslice"
	"log"
	"math"
	"time"
)

var (
	oldJwSet = make([]model.Dto, 0)
	oldSeSet = make([]model.Dto, 0)
)

func (s *Server) polling() {
	duration := time.Duration(s.Config.ServerConfig.PollingDuration) * time.Second
	moreStr := fmt.Sprintf(
		"--- \n+ For more, please visit [华工教务通知](%s) and [软院公务通知](%s)",
		s.Static.JwHomepage, s.Static.SeHomepage,
	)
	for {
		newJwSet, err1 := service.FetchJwNotice(s.Static)
		newSeSet, err2 := service.FetchSeNotice(s.Static)
		if err1 != nil || err2 != nil {
			time.Sleep(duration)
			continue
		}

		jwDiff := xslice.Its(xslice.SliceDiff(xslice.Sti(newJwSet), xslice.Sti(oldJwSet)), model.Dto{}).([]model.Dto)
		seDiff := xslice.Its(xslice.SliceDiff(xslice.Sti(newSeSet), xslice.Sti(oldSeSet)), model.Dto{}).([]model.Dto)

		sendList := make([]model.Dto, len(jwDiff)+len(seDiff))
		sendList = append(sendList, jwDiff...)
		sendList = append(sendList, seDiff...)
		s.sendDtoSlice(sendList, moreStr)

		oldJwSet = newJwSet
		oldSeSet = newSeSet
	}
}

func (s *Server) sendDtoSlice(dtos []model.Dto, tail string) {
	maxCnt := s.Static.SendMaxCount
	sendTimes := int(math.Ceil(float64(len(dtos)) / float64(maxCnt)))
	for i := 0; i < sendTimes; i++ {
		// split
		from := i * maxCnt     // 0
		to := (i + 1) * maxCnt // 10
		if to >= len(dtos) {
			to = len(dtos)
		}

		// join
		msg := ""
		for j := from; j < to; j++ {
			msg += fmt.Sprintf("+ %s\n", dtos[j].String())
		}

		// send
		if msg != "" {
			if i == sendTimes-1 { // last
				msg += tail
			}
			s.send(fmt.Sprintf("教务系统通知及软院公务通知 (第 %d 条，共 %d 条)", i+1, sendTimes), msg)
			log.Printf("Sent %d notice(s) success", to-from)
		}
	}
}
