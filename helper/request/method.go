package request

import (
	"encoding/json"
	"time"
)

type H map[string]string

// GET

func Get(url string, headers H) ([]byte, error) {

	c := Client{"GET", url, "", headers, 0}
	return c.Request()

}

func TextGet(url string, headers H) (string, error) {

	c := Client{"GET", url, "", headers, 0}
	return c.TextRequest()

}

func TimingGet(url string, headers H, timeout int64) string {

	c := Client{"GET", url, "", headers, time.Duration(timeout) * time.Second}

	if res, err := c.TextRequest(); err == nil && res != "" {
		return res
	}

	return ""

}

// POST

func Post(url, query string, headers H) ([]byte, error) {

	c := Client{"POST", url, query, headers, 0}
	return c.Request()

}

func JsonPost(url string, query any, headers H) ([]byte, error) {

	data, err := json.Marshal(query)

	if err != nil {
		return nil, err
	}

	c := Client{"POST", url, string(data), headers, 0}
	return c.JsonRequest()

}

func TextPost(url string, query string, headers H) (string, error) {

	c := Client{"POST", url, query, headers, 0}
	return c.TextRequest()

}

// PUT

func Put(url, query string, headers H) ([]byte, error) {

	c := Client{"PUT", url, query, headers, 0}
	return c.Request()

}

func JsonPut(url string, query any, headers H) ([]byte, error) {

	data, err := json.Marshal(query)

	if err != nil {
		return nil, err
	}

	c := Client{"PUT", url, string(data), headers, 0}
	return c.JsonRequest()

}

// PATCH

func Patch(url, query string, headers H) ([]byte, error) {

	c := Client{"PATCH", url, query, headers, 0}
	return c.Request()

}

func JsonPatch(url string, query any, headers H) ([]byte, error) {

	data, err := json.Marshal(query)

	if err != nil {
		return nil, err
	}

	c := Client{"PATCH", url, string(data), headers, 0}
	return c.JsonRequest()

}

// DELETE

func Delete(url string, headers H) ([]byte, error) {

	c := Client{"DELETE", url, "", headers, 0}

	return c.Request()

}

func JsonDelete(url string, headers H) ([]byte, error) {

	c := Client{"DELETE", url, "", headers, 0}

	return c.JsonRequest()

}
