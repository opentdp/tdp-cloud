package cloudflare

import (
	"encoding/json"
)

var endpoint = "https://api.cloudflare.com/client/v4"

type Params struct {
	Token   string `note:"Api Token"`
	Method  string `binding:"required"`
	Path    string `binding:"required"`
	Query   string `note:"请求参数"`
	Payload json.RawMessage
}

type Response struct {
	Success    bool
	Errors     []ResponseInfo
	Messages   []ResponseInfo
	Result     any
	ResultInfo ResultInfo
}

type ResponseInfo struct {
	Code    int
	Message string
}

type ResultInfo struct {
	Page       int
	PerPage    int
	TotalPages int
	Count      int
	Total      int
	Cursor     string
	Cursors    ResultInfoCursors
}

type ResultInfoCursors struct {
	Before string
	After  string
}

////

type OutputResult struct {
	Result   any
	Messages string `json:",omitempty"`
	DataInfo struct {
		Page    int `json:",omitempty"`
		PerPage int `json:",omitempty"`
		Total   int `json:",omitempty"`
	}
}
