package service

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func httpGet(url string) ([]byte, *http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode != 200 {
		return nil, nil, errors.New("service: get non-200 response")
	}

	body := resp.Body
	defer body.Close()
	bs, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, nil, err
	}

	return bs, resp, nil
}
