package service

import (
	"encoding/json"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xerror"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/model"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/config"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/static"
	"log"
	"strings"
	"sync"
	"time"
)

func getJwNotices() ([]*model.NoticeItem, error) {
	bs, _, err := httpGet(static.JwNoticeApi)
	if err != nil {
		return nil, err
	}

	result := &model.NoticeItemResult{}
	err = json.Unmarshal(bs, result)
	if err != nil {
		return nil, err
	}
	return result.Data.Data, nil
}

func getSeNotices() ([]*model.NoticeItem, error) {
	bs, _, err := httpGet(static.SeNoticeApi)
	if err != nil {
		return nil, err
	}

	result := &model.NoticeItemResult{}
	err = json.Unmarshal(bs, result)
	if err != nil {
		return nil, err
	}
	return result.Data.Data, nil
}

func getGrNotices() ([]*model.NoticeItem, error) {
	bs, _, err := httpGet(static.GrNoticeApi)
	if err != nil {
		return nil, err
	}

	result := &model.NoticeItemResult{}
	err = json.Unmarshal(bs, result)
	if err != nil {
		return nil, err
	}
	return result.Data.Data, nil
}

func getGzicNotices() ([]*model.NoticeItem, error) {
	bs, _, err := httpGet(static.GzicNoticeApi)
	if err != nil {
		return nil, err
	}

	result := &model.NoticeItemResult{}
	err = json.Unmarshal(bs, result)
	if err != nil {
		return nil, err
	}
	return result.Data.Data, nil
}

func GetNoticeItems() ([]*model.NoticeItem, error) {
	functions := []func() ([]*model.NoticeItem, error){getJwNotices, getSeNotices, getGrNotices, getGzicNotices}
	out := make([]*model.NoticeItem, 0)
	errs := make([]error, 0, len(functions))
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for _, f := range functions {
		f := f
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Printf("%p", f)
			result, err := f()
			log.Println(len(result))
			mu.Lock()
			defer mu.Unlock()

			for _, item := range result {
				if timeInRange(item.Date, config.Configs().Task.NotifierTimeRange) {
					out = append(out, item)
				}
			}
			errs = append(errs, err)
		}()
	}
	wg.Wait()

	if len(out) == 0 {
		return nil, xerror.Combine(errs...)
	}
	model.SortNoticeItemSlice(out)
	return out, nil
}

func timeInRange(ymd string, dayRange int32) bool {
	give, err := time.ParseInLocation("2006-01-02", ymd, time.Local)
	if err != nil {
		log.Printf("Warning: failed to parse time `%s`: `%v`", ymd, err)
		return false
	}
	du := time.Duration(dayRange) * 24 * time.Hour // unit: day
	return give.After(time.Now().Add(-du))
}

func RenderNoticeItems(items []*model.NoticeItem, fromTask bool) string {
	if len(items) == 0 {
		return ""
	}

	sb := strings.Builder{}
	if fromTask {
		sb.WriteString("*华工发布新通知")
	} else {
		sb.WriteString("*华工最新通知内容")
	}
	sb.WriteString(" (内容仅来自教务处、软院、研究生院、GZIC)*\n=====\n")

	for idx, item := range items {
		s := fmt.Sprintf("%s: [%s](%s) - %s", item.Type, item.Title, item.Url, item.Date)
		sb.WriteString(fmt.Sprintf("%d. %s\n", idx+1, s))
	}

	footer := fmt.Sprintf("=====\n更多信息，请查阅 [华工教务处通知公告](%s)、[软件学院新闻资讯](%s)、[研究生院通知公告](%s)、[华工 GZIC 通知公告](%s)。",
		static.JwNoticeHomepage, static.SeNoticeHomepage, static.GrNoticeHomepage, static.GzicNoticeHomepage)
	sb.WriteString(footer)
	return sb.String()
}

func GetNoticeLinks() string {
	return fmt.Sprintf("1. [华工教务处通知公告](%s)\n2. [软件学院新闻资讯](%s)\n3. [研究生院通知公告](%s)\n4. [华工 GZIC 通知公告](%s)",
		static.JwNoticeHomepage, static.SeNoticeHomepage, static.GrNoticeHomepage, static.GzicNoticeHomepage)
}
