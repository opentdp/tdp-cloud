package request

import (
	"net/http"
)

func Get(url string, headers map[string]string) (string, error) {

	if req, err := http.NewRequest("GET", url, nil); err == nil {
		body, err := Client(req, headers)
		return string(body), err
	} else {
		return "", err
	}

}

func GetJson(url string, headers map[string]string) ([]byte, error) {

	if req, err := http.NewRequest("GET", url, nil); err == nil {
		return Client(req, headers)
	} else {
		return nil, err
	}

}
