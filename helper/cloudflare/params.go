package cloudflare

var endpoint = "https://api.cloudflare.com/client/v4"

type Params struct {
	Token   string `note:"Api Token"`
	Uri     string `binding:"required"`
	Query   string `note:"请求参数"`
	Payload any    `note:"结构化数据"`
}

func (rq *Params) GetUrl() string {

	return endpoint + rq.Uri + "?" + rq.Query

}

func (rq *Params) GetHeader() map[string]string {

	return map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + rq.Token,
	}

}
