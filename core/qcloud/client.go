package qcloud

import (
	"encoding/json"
	"os"

	tc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	th "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	tp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

type Params struct {
	Service   string
	Version   string
	Action    string
	Payload   []byte
	Region    string
	SecretId  string
	SecretKey string
}

type Response struct {
	Response interface{} `json:"Response"`
}

func NewRequest(rp *Params) (res *Response, err error) {

	request := th.NewCommonRequest(rp.Service, rp.Version, rp.Action)

	if rp.Payload != nil {
		request.SetActionParameters(rp.Payload)
	}

	client := NewClient(rp)
	response := th.NewCommonResponse()

	if err = client.Send(request, response); err != nil {
		return
	}

	res = &Response{}
	body := response.GetBody()

	if err = json.Unmarshal(body, res); err != nil {
		return
	}

	return

}

func NewClient(rp *Params) (c *tc.Client) {

	profile := tp.NewClientProfile()

	if os.Getenv("IS_DEBUG") != "" {
		profile.Debug = true
	}

	if rp.Region != "" {
		profile.HttpProfile.Endpoint = rp.Service + "." + rp.Region + ".tencentcloudapi.com"
	}

	credential := tc.NewCredential(rp.SecretId, rp.SecretKey)

	c = tc.NewCommonClient(credential, rp.Region, profile)

	return

}
