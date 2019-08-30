package models

type RetJson struct {
	List    []ItemJson `json:"list"`
	Message string     `json:"message"`
	Pagenum int        `json:"pagenum"`
	Row     int        `json:"row"`
	Success bool       `json:"success"`
	Total   int        `json:"total"`
}

type ItemJson struct {
	CreateTime string `json:"createTime"`
	Id         string `json:"id"`
	IsNew      bool   `json:"isNew"`
	Tag        int    `json:"tag"`
	Title      string `json:"title"`
}
