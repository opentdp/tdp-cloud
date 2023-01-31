package cloudflare

import (
	"encoding/json"
	"errors"

	"tdp-cloud/helper/request"
)

func Request(rp *Params) (*OutputResult, error) {

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

func parseBody(body []byte) (*OutputResult, error) {

	res := &Response{}
	err := json.Unmarshal(body, res)

	if err != nil {
		return nil, err
	}

	out := &OutputResult{
		Result: res.Result,
	}

	if cap(res.Messages) > 0 {
		out.Messages = (res.Messages[0]).Message
	}

	if cap(res.Errors) > 0 {
		err = errors.New((res.Errors[0]).Message)
	}

	if res.ResultInfo.PerPage > 0 {
		out.DataInfo.Page = res.ResultInfo.Page
		out.DataInfo.PerPage = res.ResultInfo.PerPage
		out.DataInfo.Total = res.ResultInfo.Total
	}

	return out, err

}
