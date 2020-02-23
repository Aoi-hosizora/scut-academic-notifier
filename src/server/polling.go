package server

import (
	"fmt"
	"github.com/Aoi-hosizora/SCUT_Academic_Notifier/src/model"
	"github.com/Aoi-hosizora/SCUT_Academic_Notifier/src/service"
	"github.com/Aoi-hosizora/SCUT_Academic_Notifier/src/util"
	"github.com/Aoi-hosizora/ahlib/xcondition"
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
			log.Printf("Failed to fetch notice, jw: %v, se: %v", err1, err2)
			time.Sleep(duration)
			continue
		}

		jwDiff := xslice.Its(xslice.SliceDiff(xslice.Sti(newJwSet), xslice.Sti(oldJwSet)), model.Dto{}).([]model.Dto)
		seDiff := xslice.Its(xslice.SliceDiff(xslice.Sti(newSeSet), xslice.Sti(oldSeSet)), model.Dto{}).([]model.Dto)

		allDiffList := make([]model.Dto, 0)
		allDiffList = append(allDiffList, jwDiff...)
		allDiffList = append(allDiffList, seDiff...)

		if len(allDiffList) == 0 {
			log.Println("Polling once and not found new notice")
			time.Sleep(duration)
			continue
		}

		// filter
		sendList := make([]model.Dto, 0)
		for _, item := range allDiffList {
			if util.IsTimeInRange(item.Date, s.Static.SendRange) {
				sendList = append(sendList, item)
			}
		}

		s.sendDtoSlice(sendList, moreStr)
		oldJwSet = newJwSet
		oldSeSet = newSeSet

		time.Sleep(duration)
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
			title := fmt.Sprintf("教务系统通知 (第 %d 条，共 %d 条)", i+1, sendTimes)
			ok := s.send(title, msg)
			log.Printf("Sent %d(%d ~ %d in %d) notice(s) %s",
				to-from, from, to, len(dtos),
				xcondition.IfThenElse(ok, "success", "failed").(string),
			)
		}
	}
}
