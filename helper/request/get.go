package request

import (
	"net/http"
)

func Get(url string, headers map[string]string) (string, error) {

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	return TextClient(req, headers)

}

func GetJson(url string, headers map[string]string) ([]byte, error) {

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	return Client(req, headers)

}
