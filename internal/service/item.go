package service

import (
	"encoding/json"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xerror"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/model"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/config"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/static"
	"log"
	"strings"
	"time"
)

func GetJwItems() ([]*model.PostItem, error) {
	bs, _, err := httpGet(static.JwApi)
	if err != nil {
		return nil, err
	}

	result := &model.PostItemDto{}
	err = json.Unmarshal(bs, result)
	if err != nil {
		return nil, err
	}
	return result.Data.Data, nil
}

func GetSeItems() ([]*model.PostItem, error) {
	bs, _, err := httpGet(static.SeApi)
	if err != nil {
		return nil, err
	}

	result := &model.PostItemDto{}
	err = json.Unmarshal(bs, result)
	if err != nil {
		return nil, err
	}
	return result.Data.Data, nil
}

func GetNoticeItems() ([]*model.PostItem, error) {
	jwItems, err1 := GetJwItems()
	seItems, err2 := GetSeItems()

	out := make([]*model.PostItem, 0)
	for _, item := range jwItems {
		if TimeInRange(item.Date, config.Configs().Task.NotifierTimeRange) {
			out = append(out, item)
		}
	}
	for _, item := range seItems {
		if TimeInRange(item.Date, config.Configs().Task.NotifierTimeRange) {
			out = append(out, item)
		}
	}

	if len(out) == 0 {
		return nil, xerror.Combine(err1, err2)
	}
	return out, nil
}

func TimeInRange(ymd string, dayRange int32) bool {
	give, err := time.ParseInLocation("2006-01-02", ymd, time.Local)
	if err != nil {
		log.Printf("Warning: failed to parse time `%s`: `%v`", ymd, err)
		return false
	}
	du := time.Duration(dayRange) * 24 * time.Hour // unit: day
	return give.After(time.Now().Add(-du))
}

func RenderNoticeItems(items []*model.PostItem, fromTask bool) string {
	if len(items) == 0 {
		return ""
	}
	sb := strings.Builder{}
	if fromTask {
		sb.WriteString("*华工发布新通知 (仅来源于教务处、软院)*\n=====\n")
	} else {
		sb.WriteString("*华工最新通知列表 (仅来源于教务处、软院)*\n=====\n")
	}
	for idx, item := range items {
		s := fmt.Sprintf("%s: [%s](%s) - %s", item.Type, item.Title, item.Url, item.Date)
		sb.WriteString(fmt.Sprintf("%d. %s\n", idx+1, s))
	}
	footer := fmt.Sprintf("=====\n更多信息，请查阅 [华工教务处通知](%s) 以及 [华工软院公务通知](%s)。", static.JwHomepage, static.SeHomepage)
	sb.WriteString(footer)
	return sb.String()
}
