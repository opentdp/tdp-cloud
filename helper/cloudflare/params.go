package cloudflare

import "encoding/json"

var endpoint = "https://api.cloudflare.com/client/v4"

type Params struct {
	Token   string `note:"Api Token"`
	Method  string `binding:"required"`
	Path    string `binding:"required"`
	Query   string `note:"请求参数"`
	Payload json.RawMessage
}
