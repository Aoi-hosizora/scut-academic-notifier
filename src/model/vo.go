package model

type JwVo struct {
	List []struct {
		CreateTime string `json:"createTime"`
		Id         string `json:"id"`
		IsNew      bool   `json:"isNew"`
		Tag        int    `json:"tag"`
		Title      string `json:"title"`
	} `json:"list"`
	Message string `json:"message"`
	PageNum int    `json:"pagenum"`
	Row     int    `json:"row"`
	Success bool   `json:"success"`
	Total   int    `json:"total"`
}

type SeVo struct {
	TagIdx  int
	Content string
}
