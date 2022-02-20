package service

import (
	"context"
	"fmt"
	"io"
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
		return nil, nil, fmt.Errorf("service: get non-200 response when requesting `%s`", req.URL.String())
	}

	body := resp.Body
	defer body.Close()
	bs, err := io.ReadAll(body)
	if err != nil {
		return nil, nil, err
	}

	return bs, resp, nil
}
