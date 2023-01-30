package request

import (
	"io"
	"net/http"
)

func Client(req *http.Request, headers map[string]string) ([]byte, error) {

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return io.ReadAll(resp.Body)

}

func TextClient(req *http.Request, headers map[string]string) (string, error) {

	body, err := Client(req, headers)
	return string(body), err

}
