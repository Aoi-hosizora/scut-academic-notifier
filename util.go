package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/fatih/set.v0"
)

type SckeyConfig struct {
	SCKEY string
}

func getConfig(url string) *SckeyConfig {
	// 返回 Json Config

	cfg := &SckeyConfig{}
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

func getTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func toMarkdown(msg string) string {
	// MarkDown 增加项目符号
	return "+ " + strings.Replace(msg, "\n", "\n+ ", -1)
}

func putNotifier(Sckey string, title string, msg string) {
	// 通过 Server 酱发送信息

	msg = fmt.Sprintf("%s (%s)", toMarkdown(msg), getTimeString())

	// url.QueryEscape 转化 url
	url := fmt.Sprintf("https://sc.ftqq.com/%s.send?text=%s&desp=%s", Sckey, url.QueryEscape(title), url.QueryEscape(msg))

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != 200 {
		panic(res.Status)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func getSet(l ListArrAndTop) set.Interface {
	// 将 ListArr 转换成 Set
	s := set.New(set.ThreadSafe)
	for i := 0; i < NORMAL_NEWS_CNT+1; i++ {
		for j := l[i].Front(); j != nil; j = j.Next() {
			if i == 0 {
				s.Add(j.Value.(*SCUT_TopNews).String())
			} else {
				s.Add(j.Value.(*SCUT_NormalNews).String())
			}
		}
	}
	return s
}

func getSetStr(s set.Interface) string {
	strret := ""
	s.Each(func(i interface{}) bool {
		strret = strret + fmt.Sprintln(i)
		return true
	})
	return strret
}
