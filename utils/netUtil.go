package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Aoi-hosizora/Academic_Notifier/models"
	"github.com/PuerkitoBio/goquery"
)

// Server Chan Url: `Sckey` `title` `msg`
var ServerChanUrl string = "https://sc.ftqq.com/%s.send?text=%s&desp=%s"

// 教务通知项目页 `id`
// var JWViewUrl string = "http://jw.scut.edu.cn/zhinan/cms/article/view.do?type=posts&id=%s"
var JWViewUrl string = "http://jw.scut.edu.cn/dist/#/detail/index?id=%s&type=notice"

// 软件学院新闻资讯页 `sse`
var SeViewUrl string = "http://www2.scut.edu.cn/%s"

// 通过 Server 酱发送信息 (POST)
func SendNotifier(Sckey string, title string, msg string) {

	// 将发送内容加上时间
	msg = fmt.Sprintf("> At %s Send: \n\n%s", GetNowTimeString(), msg)

	// url.QueryEscape 转化 url
	url := fmt.Sprintf(ServerChanUrl, Sckey, url.QueryEscape(title), url.QueryEscape(msg))
	res, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader("name=cjb"))

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
	log.Println(string(body))
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 教务通知

// Post 获得 Json
func GetPostData(postUrl string, tag int, pageSize int) string {
	client := &http.Client{}

	form := url.Values{}
	form.Add("tag", strconv.Itoa(tag))
	form.Add("pageNum", "1")
	form.Add("pageSize", strconv.Itoa(pageSize))
	form.Add("keyword", "")

	req, err := http.NewRequest("POST", postUrl, strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Referer", "http://jw.scut.edu.cn/zhinan/cms/toPosts.do")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return string(body)
}

// 处理 Json 字符串 -> NoticeItem
func ParseJson(Json string) []models.NoticeItem {
	var s models.RetJson
	err := json.Unmarshal([]byte(Json), &s)
	if err != nil {
		panic(err)
	}

	l := make([]models.NoticeItem, len(s.List))
	for i := 0; i < len(s.List); i++ {
		nt := s.List[i]
		notice := models.NoticeItem{
			Title: nt.Title,
			Url:   fmt.Sprintf(JWViewUrl, nt.Id),
			Type:  GetTypeFromTag(nt.Tag),
			Date:  nt.CreateTime,
			IsNew: nt.IsNew,
		}
		l[i] = notice
	}
	return l
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 学院通知

type pSel = *goquery.Selection
type pDoc = *goquery.Document

func GetSENotices(SESchoolUrl string, PartUrls []string, PartNames []string) []models.NoticeItem {
	var rets = []models.NoticeItem{}
	for k, v := range PartUrls {
		ls := _GetSEPartNotices(fmt.Sprintf(SESchoolUrl, v), "软院 - "+PartNames[k])
		if ls == nil {
			return nil
		}
		for _, l := range ls {
			rets = append(rets, l)
		}
	}
	return rets
}

// 解析 html 返回通知
func _GetSEPartNotices(SESchoolUrl string, PartName string) []models.NoticeItem {
	client := &http.Client{}
	resp, err := client.Get(SESchoolUrl)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil
	}

	return _ParseSENotices(doc, PartName)
}

// 解析 html 内容
func _ParseSENotices(doc pDoc, PartName string) []models.NoticeItem {
	lis := doc.Find("ul.news_ul > li.news_li")
	l := make([]models.NoticeItem, lis.Size())
	lis.Each(func(i int, s pSel) {
		a := s.Find(".news_title a")
		meta := s.Find(".news_title span.news_meta")
		notice := models.NoticeItem{
			Title: a.Text(),
			Url:   fmt.Sprintf(SeViewUrl, a.AttrOr("href", "")),
			Type:  PartName,
			Date:  meta.Text(),
			IsNew: true,
		}
		l[i] = notice
	})
	return l
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 其他处理

// a1 + a2
func ToArrayAdd(a1 []models.NoticeItem, a2 []models.NoticeItem) []models.NoticeItem {
	for _, k := range a2 {
		a1 = append(a1, k)
	}
	return a1
}

// [is / isnot new] - [old]
func ToArrayDifference(new []models.NoticeItem, old []models.NoticeItem) []models.NoticeItem {
	diff := make([]models.NoticeItem, len(new))
	num := 0
	for i := 0; i < len(new); i++ {
		if new[i].IsNew {
			has := false
			for _, v := range old {
				if v.Title == new[i].Title {
					has = true
				}
			}
			if !has {
				diff[num] = new[i]
				num += 1
			}
		}
	}

	ret := make([]models.NoticeItem, num)
	for i := 0; i < len(ret); i++ {
		ret[i] = diff[i]
	}
	return ret
}

// Api tag -> Type str
func GetTypeFromTag(tag int) string {
	switch tag {
	case 1:
		return "选课"
	case 2:
		return "考试"
	case 3:
		return "实践"
	case 4:
		return "交流"
	case 5:
		return "教师"
	case 6:
		return "信息"
	}
	return fmt.Sprintf("未知-%d", tag)
}
