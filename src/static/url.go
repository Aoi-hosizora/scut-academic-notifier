package static

var (
	JwHomepage = "http://jw.scut.edu.cn/zhinan/cms/index.do"
	JwReferer  = "http://jw.scut.edu.cn/zhinan/cms/toPosts.do"
	JwApiUrl   = "http://jw.scut.edu.cn/zhinan/cms/article/v2/findInformNotice.do"
	JwTagNames = []string{"选课", "考试", "实践", "交流", "教师", "信息"}
	JwItemUrl  = "http://jw.scut.edu.cn/dist/#/detail/index?id=%s&type=notice"
	// JwItemUrl  = "http://jw.scut.edu.cn/zhinan/cms/article/view.do?type=posts&id=%s"

	SeHomepage    = "http://www2.scut.edu.cn/sse/xyjd_17232/list.htm"
	SeWebUrl      = "http://www2.scut.edu.cn/sse/%s/list.htm"
	SeWebUrlParts = []string{"xyjd_17232", "17235", "17236", "gwtz", "kytz"}
	SeTagNames    = []string{"学院焦点", "本科生通知", "研究生通知", "公务通知", "科研通知"}
	SeItemUrl     = "http://www2.scut.edu.cn/%s"
)
