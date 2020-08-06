package service

import (
	"encoding/json"
	"fmt"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/config"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/model"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func FetchJwNotice(static *config.Static) ([]model.Dto, error) {
	vo, err := postJw(static)
	if err != nil {
		return nil, err
	}
	log.Println("Fetch jw notice success")

	ret := make([]model.Dto, len(vo.List))
	for i, vo := range vo.List {
		ret[i] = model.Dto{
			Title: vo.Title,
			Url:   fmt.Sprintf(static.JwItemUrl, vo.Id),
			Type:  static.JwTagNames[vo.Tag-1],
			Date:  strings.ReplaceAll(vo.CreateTime, ".", "-"), // 2020.01.01
			IsNew: vo.IsNew,
		}
	}
	return ret, nil
}

func postJw(static *config.Static) (*model.JwVo, error) {
	client := &http.Client{}

	// tag === 0
	form := &url.Values{}
	form.Add("tag", "0")
	form.Add("pageNum", "1")
	form.Add("pageSize", "50")
	form.Add("keyword", "")
	formBody := strings.NewReader(form.Encode())

	req, err := http.NewRequest("POST", static.JwApiUrl, formBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Referer", static.JwReferer)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	vo := &model.JwVo{}
	err = json.Unmarshal(body, &vo)
	if err != nil {
		return nil, err
	}
	return vo, nil
}
