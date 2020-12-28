package model

import "fmt"

type PostItem struct {
	Title string `json:"title"`
	Url   string `json:"url"`
	Type  string `json:"type"`
	Date  string `json:"date"`
}

func (d *PostItem) String() string {
	return fmt.Sprintf("%s: [%s](%s) - %s", d.Type, d.Title, d.Url, d.Date)
}

func ItemSliceDiff(s1 []*PostItem, s2 []*PostItem) []*PostItem {
	result := make([]*PostItem, 0)
	for _, item1 := range s1 {
		exist := false
		for _, item2 := range s2 {
			if item1.Type == item2.Type && item1.Title == item2.Title {
				exist = true
				break
			}
		}
		if !exist {
			result = append(result, item1)
		}
	}
	return result
}
