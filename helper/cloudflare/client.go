package cloudflare

import (
	"encoding/json"
	"errors"

	"tdp-cloud/helper/request"
)

func Request(rq *ReqeustParam) (any, error) {

	client := request.Client{
		Method: rq.Method,
		Url:    endpoint + rq.Path,
		Data:   string(rq.Payload),
		Headers: request.H{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + rq.Token,
		},
	}

	if rq.Query != "" {
		client.Url += "?" + rq.Query
	}

	body, err := client.JsonRequest()

	if err != nil {
		return nil, err
	}

	return parseBody(body)

}

func parseBody(body []byte) (any, error) {

	res := &ResponseData{}
	err := json.Unmarshal(body, res)

	if err != nil {
		return nil, err
	}

	out := &Output{res.Result, ""}

	if cap(res.Messages) > 0 {
		out.Messages = (res.Messages[0]).Message
	}

	if cap(res.Errors) > 0 {
		err = errors.New((res.Errors[0]).Message)
	}

	if res.ResultInfo.PerPage == 0 {
		return out, err
	}

	inf := &OutputInfo{
		res.ResultInfo.Page, res.ResultInfo.PerPage, res.ResultInfo.Total,
	}

	return &OutputWithInfo{out, inf}, err

}
