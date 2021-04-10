package service

import (
	"io/ioutil"
	"net/http"
)

const (
	JwApi = "http://api.common.aoihosizora.top/scut/jw"
	SeApi = "http://api.common.aoihosizora.top/scut/se"

	JwHomepage = "http://jw.scut.edu.cn/zhinan/cms/index.do"
	SeHomepage = "http://www2.scut.edu.cn/sse/xyjd_17232/list.htm"
)

func httpGet(url string) ([]byte, *http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	body := resp.Body
	defer body.Close()
	bs, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, nil, err
	}

	return bs, resp, nil
}
