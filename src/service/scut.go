package service

import (
	"encoding/json"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/model"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/static"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func httpGet(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body := resp.Body
	defer body.Close()
	bs, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

type r struct {
	Data struct {
		Data []*model.PostItem `json:"data"`
	} `json:"data"`
}

func GetJwItems() ([]*model.PostItem, error) {
	bs, err := httpGet(static.JwApi)
	if err != nil {
		return nil, err
	}
	result := &r{}
	err = json.Unmarshal(bs, result)
	if err != nil {
		return nil, err
	}
	return result.Data.Data, nil
}

func GetSeItems() ([]*model.PostItem, error) {
	bs, err := httpGet(static.SeApi)
	if err != nil {
		return nil, err
	}
	result := &r{}
	err = json.Unmarshal(bs, result)
	if err != nil {
		return nil, err
	}
	return result.Data.Data, nil
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
