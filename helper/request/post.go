package request

import (
	"net/http"
	"strings"
)

func Post(url, query string, headers map[string]string) (string, error) {

	rd := strings.NewReader(query)

	if headers["Content-Type"] == "" {
		headers["Content-Type"] = "application/x-www-form-urlencoded; charset=UTF-8"
	}

	if req, err := http.NewRequest("POST", url, rd); err == nil {
		body, err := Client(req, headers)
		return string(body), err
	} else {
		return "", err
	}

}

func PostJson(url string, query []byte, headers map[string]string) ([]byte, error) {

	rd := strings.NewReader(string(query))

	if headers["Content-Type"] == "" {
		headers["Content-Type"] = "application/json"
	}

	if req, err := http.NewRequest("POST", url, rd); err == nil {
		return Client(req, headers)
	} else {
		return nil, err
	}

}
