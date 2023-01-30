package cloudflare

import (
	"encoding/json"

	"tdp-cloud/helper/request"
)

type Params struct {
	ApiToken string
	Uri      string `binding:"required"`
	Query    string
	Payload  any
}

type Response struct {
	Result     any
	Success    bool
	Errors     any
	Messages   any
	ResultInfo any
}

var Endpoint = "https://api.cloudflare.com/client/v4"

func Get(rq *Params) (*Response, error) {

	res := &Response{}

	url := Endpoint + rq.Uri + "?" + rq.Query

	header := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + rq.ApiToken,
	}

	body, err := request.GetJson(url, header)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(body, res)
	return res, err

}

func Post(rq *Params) (*Response, error) {

	res := &Response{}

	url := Endpoint + rq.Uri + "?" + rq.Query

	header := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + rq.ApiToken,
	}

	data, err := json.Marshal(rq.Payload)
	if err != nil {
		return res, err
	}

	body, err := request.PostJson(url, data, header)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(body, res)
	return res, err

}
