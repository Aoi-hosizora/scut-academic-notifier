package service

import (
	"fmt"
	"github.com/Aoi-hosizora/SCUT_Academic_Notifier/src/config"
	"github.com/Aoi-hosizora/SCUT_Academic_Notifier/src/model"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"strings"
)

func FetchSeNotice(static *config.Static) ([]model.Dto, error) {
	vos := make([]*model.SeVo, len(static.SeWebUrlParts))
	for k := range static.SeWebUrlParts {
		vo, err := getSe(static, k)
		if err != nil {
			return nil, err
		}
		vos[k] = vo
	}

	ret := make([]model.Dto, 0)
	for _, vo := range vos {
		dtos, err := parseSe(static, vo)
		if err != nil {
			return nil, err
		}
		ret = append(ret, dtos...)
	}
	return ret, nil
}

func getSe(static *config.Static, tagIdx int) (*model.SeVo, error) {
	client := &http.Client{}

	url := fmt.Sprintf(static.SeWebUrl, static.SeWebUrlParts[tagIdx])
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &model.SeVo{
		TagIdx:  tagIdx,
		Content: string(body),
	}, nil
}

func parseSe(static *config.Static, vo *model.SeVo) ([]model.Dto, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(vo.Content))
	if err != nil {
		return nil, err
	}
	lis := doc.Find("ul.news_ul > li.news_li")
	ret := make([]model.Dto, lis.Size())

	lis.Each(func(i int, s *goquery.Selection) {
		a := s.Find(".news_title a")
		meta := s.Find("span.news_meta")

		ret[i] = model.Dto{
			Title: a.Text(),
			Url:   fmt.Sprintf(static.SeItemUrl, a.AttrOr("href", "")),
			Type:  static.SeTagNames[vo.TagIdx],
			Date:  meta.Text(), // 2019-10-01
			IsNew: true,
		}
	})
	return ret, nil
}
