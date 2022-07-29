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

func NewRequest(rp *Params) (*Response, error) {

	request := th.NewCommonRequest(rp.Service, rp.Version, rp.Action)

	if err := request.SetActionParameters(rp.Payload); err != nil {
		return nil, err
	}

	response := th.NewCommonResponse()

	if err := NewClient(rp).Send(request, response); err != nil {
		return nil, err
	}

	var res = &Response{}
	if err := json.Unmarshal(response.GetBody(), res); err != nil {
		return nil, err
	}

	return res, nil

}

func NewClient(rp *Params) *tc.Client {

	profile := tp.NewClientProfile()

	if os.Getenv("IS_DEBUG") != "" {
		profile.Debug = true
	}

	if rp.Region != "" {
		profile.HttpProfile.Endpoint = rp.Service + "." + rp.Region + ".tencentcloudapi.com"
	}

	credential := tc.NewCredential(rp.SecretId, rp.SecretKey)

	return tc.NewCommonClient(credential, rp.Region, profile)

}
