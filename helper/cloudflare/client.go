package cloudflare

import (
	"encoding/json"
	"errors"

	"tdp-cloud/helper/request"
)

func Get(rq *Params) (any, error) {

	url := rq.GetUrl()
	header := rq.GetHeader()
	body, err := request.GetJson(url, header)

	if err != nil {
		return nil, err
	}

	return parseBody(body)

}

func Post(rq *Params) (any, error) {

	url := rq.GetUrl()
	header := rq.GetHeader()
	body, err := request.PostJson(url, rq.Payload, header)

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
