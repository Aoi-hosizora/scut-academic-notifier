package config

import (
	"fmt"
	"time"
)

type Static struct {
	ServerChanUrl string

	SendMaxCount int
	SendRange    time.Duration

	JwHomepage string   // jw homepage (mobile)
	JwReferer  string   // request referer
	JwApiUrl   string   // request api url
	JwTagNames []string // response tag names
	JwItemUrl  string   // response item url (mobile)

	SeHomepage    string   // se homepage
	SeWebUrl      string   // request web url
	SeWebUrlParts []string // request web url parts
	SeTagNames    []string // response tag names
	SeItemUrl     string   // response item url
}

func LoadStatic() *Static {
	static := &Static{
		ServerChanUrl: "https://sc.ftqq.com/%s.send?text=%s&desp=%s", // `Sckey` `title` `msg`

		SendMaxCount: 10,
		SendRange:    60 * 24 * time.Hour, // 2 months

		JwHomepage: "http://jw.scut.edu.cn/zhinan/cms/index.do",
		JwReferer:  "http://jw.scut.edu.cn/zhinan/cms/toPosts.do",
		JwApiUrl:   "http://jw.scut.edu.cn/zhinan/cms/article/v2/findInformNotice.do",
		JwTagNames: []string{"选课", "考试", "实践", "交流", "教师", "信息"},
		JwItemUrl:  "http://jw.scut.edu.cn/dist/#/detail/index?id=%s&type=notice",
		// http://jw.scut.edu.cn/zhinan/cms/article/view.do?type=posts&id=%s

		SeWebUrl:      "http://www2.scut.edu.cn/sse/%s/list.htm",
		SeWebUrlParts: []string{"xyjd_17232", "17235", "17236", "gwtz", "kytz"},
		SeTagNames:    []string{"学院焦点", "本科生通知", "研究生通知", "公务通知", "科研通知"},
		SeItemUrl:     "http://www2.scut.edu.cn/%s",
	}
	static.SeHomepage = fmt.Sprintf(static.SeWebUrl, static.SeWebUrlParts[0])
	return static
}
