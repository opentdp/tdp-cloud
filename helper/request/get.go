package request

import (
	"net/http"
)

func Get(url string, headers map[string]string) (string, error) {

	if req, err := http.NewRequest("GET", url, nil); err == nil {
		return Client(req, headers)
	} else {
		return "", err
	}

}

func GetJson(url string, headers map[string]string) (string, error) {

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	return Get(url, headers)

}
