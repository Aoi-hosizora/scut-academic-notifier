package service

import (
	"encoding/json"
	"errors"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/model"
	"log"
	"time"
)

func GetJwItems() ([]*model.PostItem, error) {
	bs, resp, err := httpGet(JwApi)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("service: something error when http get")
	}

	result := &model.PostItemDto{}
	err = json.Unmarshal(bs, result)
	if err != nil {
		return nil, err
	}
	return result.Data.Data, nil
}

func GetSeItems() ([]*model.PostItem, error) {
	bs, resp, err := httpGet(SeApi)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("service: something error when http get")
	}

	result := &model.PostItemDto{}
	err = json.Unmarshal(bs, result)
	if err != nil {
		return nil, err
	}
	return result.Data.Data, nil
}

func CheckInTimeRange(str string, day int32) bool {
	t, err := time.ParseInLocation("2006-01-02", str, time.Local)
	if err != nil {
		log.Println("Failed to parse time:", err)
		return true
	}
	du := time.Duration(day) * 24 * time.Hour // a month
	return t.After(time.Now().Add(-du))
}
