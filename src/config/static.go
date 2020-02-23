package config

import (
	"fmt"
)

type Static struct {
	SendMaxCount  int
	ServerChanUrl string

	JwHomepage string
	JwReferer  string
	JwApiUrl   string
	JwTagNames []string
	JwItemUrl  string

	SeHomepage    string
	SeWebUrl      string
	SeWebUrlParts []string
	SeTagNames    []string
	SeItemUrl     string
}

func LoadStatic() *Static {
	static := &Static{
		SendMaxCount:  10,
		ServerChanUrl: "https://sc.ftqq.com/%s.send?text=%s&desp=%s", // `Sckey` `title` `msg`

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
