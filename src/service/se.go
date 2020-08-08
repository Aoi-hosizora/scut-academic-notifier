package service

import (
	"fmt"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/model"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/static"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func GetSeItems() ([]*model.Item, error) {
	ses := make([]*model.SePage, len(static.SeWebUrlParts))
	for idx, part := range static.SeWebUrlParts {
		url := fmt.Sprintf(static.SeWebUrl, part)
		resp, err := HttpRequest(url, "GET", nil, false)
		if err != nil {
			return nil, err
		}
		ses[idx] = &model.SePage{
			TagPart: part,
			TagName: static.SeTagNames[idx],
			Content: string(resp),
		}
	}

	ret := make([]*model.Item, 0)
	for _, se := range ses {
		items, err := parseSeItem(se)
		if err != nil {
			return nil, err
		}
		ret = append(ret, items...)
	}

	return ret, nil
}

func parseSeItem(vo *model.SePage) ([]*model.Item, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(vo.Content))
	if err != nil {
		return nil, err
	}

	lis := doc.Find("ul.news_ul > li.news_li")
	ret := make([]*model.Item, lis.Size())

	lis.Each(func(i int, s *goquery.Selection) {
		a := s.Find(".news_title a")
		meta := s.Find("span.news_meta")

		ret[i] = &model.Item{
			Title: a.Text(),
			Url:   fmt.Sprintf(static.SeItemUrl, a.AttrOr("href", "")),
			Type:  "软院" + vo.TagName,
			Date:  meta.Text(), // 2019-10-01
		}
	})
	return ret, nil
}
