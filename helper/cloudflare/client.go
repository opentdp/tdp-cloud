package cloudflare

import (
	"encoding/json"
	"errors"

	"tdp-cloud/helper/request"
)

func Request(rp *Params) (any, error) {

	client := request.Client{
		Method: rp.Method,
		Url:    endpoint + rp.Path + "?" + rp.Query,
		Data:   string(rp.Payload),
		Headers: request.H{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + rp.Token,
		},
	}

	body, err := client.JsonRequest()

	if err != nil {
		return nil, err
	}

	return parseBody(body)

}

func parseBody(body []byte) (any, error) {

	res := &Response{}
	err := json.Unmarshal(body, res)

	if err != nil {
		return res, err
	}

	if cap(res.Errors) > 0 {
		err = errors.New((res.Errors[0]).Message)
		return res, err
	}

	if cap(res.Messages) > 0 {
		msg := (res.Messages[0]).Message
		return msg, err
	}

	return res, err

}
