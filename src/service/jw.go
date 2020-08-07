package service

import (
	"encoding/json"
	"fmt"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/model"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/static"
	"net/url"
	"strings"
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
