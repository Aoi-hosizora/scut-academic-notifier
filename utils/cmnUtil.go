package utils

import (
	"encoding/json"
	"io/ioutil"
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

////////////////////////////////////////////////////////////////////////////////////
// Private:

// MarkDown 增加项目符号
func toMarkdown(msg string) string {
	return "+ " + strings.Replace(msg, "\n", "\n+ ", -1)
}
