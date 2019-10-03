package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	SCKEY string
}

// 获得 SCKEY
func GetConfig(url string) *Config {
	cfg := &Config{}
	data, err := ioutil.ReadFile(url)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

// 当前时间格式化
func GetNowTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 将 00.00 转化成 2019-00-00
func ParseSpTimeString(s string) string {
	sp := strings.Split(s, ".") // 08.30
	ret := s
	if len(sp) == 2 {
		month, err1 := strconv.Atoi(sp[0])
		day, err2 := strconv.Atoi(sp[1])

		if err1 != nil && err2 != nil {
			nttime := time.Date(time.Now().Year(), time.Month(month), day, 0, 0, 0, 0, time.Local)
			if nttime.After(time.Now()) { // 去年
				ret = fmt.Sprintf("%d-%d-%d", time.Now().Year()-1, month, day)
			} else { // 今年
				ret = fmt.Sprintf("%d-%d-%d", time.Now().Year(), month, day)
			}
		}
	}
	return ret
}
