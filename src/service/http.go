package service

import (
	"github.com/Aoi-hosizora/scut-academic-notifier/src/static"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	UserAgent   = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36"
	ContentType = "application/x-www-form-urlencoded;charset=UTF-8"
)

func HttpRequest(url string, method string, b io.Reader, useReferer bool) ([]byte, error) {
	method = strings.ToUpper(method)
	req, err := http.NewRequest(method, url, b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Content-Type", ContentType)
	if useReferer {
		req.Header.Set("Referer", static.JwReferer)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body := resp.Body
	defer body.Close()
	bs, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	return bs, nil
}
