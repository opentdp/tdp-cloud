package request

import (
	"io"
	"net/http"
)

func Client(req *http.Request, headers map[string]string) (string, error) {

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	return string(body), err

}
