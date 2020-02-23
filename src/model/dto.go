package model

import "fmt"

type Dto struct {
	Title string
	Url   string
	Type  string
	Date  string
	IsNew bool
}

func (d *Dto) String() string {
	return fmt.Sprintf("%s：[%s](%s) - %s", d.Type, d.Title, d.Url, d.Date)
}
