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
		return Client(req, headers)
	} else {
		return "", err
	}

}

func PostJson(url, query string, headers map[string]string) (string, error) {

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	return Post(url, query, headers)

}
