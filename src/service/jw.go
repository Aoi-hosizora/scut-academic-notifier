package service

import (
	"encoding/json"
	"fmt"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/model"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/static"
	"log"
	"net/url"
	"strings"
	"time"
)

func GetJwItems() ([]*model.Item, error) {
	form := &url.Values{}
	form.Add("tag", "0")
	form.Add("pageNum", "1")
	form.Add("pageSize", "50")
	form.Add("keyword", "")
	formBody := strings.NewReader(form.Encode())
	resp, err := HttpRequest(static.JwApiUrl, "POST", formBody, true)
	if err != nil {
		return nil, err
	}

	jw := &model.JwItem{}
	err = json.Unmarshal(resp, &jw)
	if err != nil {
		return nil, err
	}

	ret := make([]*model.Item, len(jw.List))
	for i, vo := range jw.List {
		ret[i] = &model.Item{
			Title: vo.Title,
			Url:   fmt.Sprintf(static.JwItemUrl, vo.Id),
			Type:  static.JwTagNames[vo.Tag-1],
			Date:  strings.ReplaceAll(vo.CreateTime, ".", "-"), // 2020.01.01
		}
	}

	return ret, nil
}

func CheckTime(str string, day int32) bool {
	t, err := time.ParseInLocation("2006-01-02", str, time.Local)
	if err != nil {
		log.Println("Failed to parse time:", err)
		return true
	}
	du := time.Duration(day) * 24 * time.Hour // a month
	return t.After(time.Now().Add(-du))
}
