package request

import (
	"encoding/json"
	"net/http"
	"strings"
)

func Post(url, query string, headers map[string]string) (string, error) {

	rd := strings.NewReader(query)
	req, err := http.NewRequest("POST", url, rd)

	if err != nil {
		return "", err
	}

	if headers["Content-Type"] == "" {
		headers["Content-Type"] = "application/x-www-form-urlencoded; charset=UTF-8"
	}

	return TextClient(req, headers)

}

func PostJson(url string, query any, headers map[string]string) ([]byte, error) {

	data, err := json.Marshal(query)

	if err != nil {
		return nil, err
	}

	rd := strings.NewReader(string(data))
	req, err := http.NewRequest("POST", url, rd)

	if err != nil {
		return nil, err
	}

	if headers["Content-Type"] == "" {
		headers["Content-Type"] = "application/json"
	}

	return Client(req, headers)

}
