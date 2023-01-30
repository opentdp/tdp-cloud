package qcloud

import (
	"encoding/json"
	"os"

	tc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	th "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	tp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func Request(rp *Params) (*Response, error) {

	request := th.NewCommonRequest(rp.Service, rp.Version, rp.Action)

	if rp.Payload != nil {
		request.SetActionParameters(rp.Payload)
	}

	client := newClient(rp)
	response := th.NewCommonResponse()

	if err := client.Send(request, response); err != nil {
		return nil, err
	}

	res := &Response{}
	body := response.GetBody()

	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil

}

func newClient(rp *Params) *tc.Client {

	cpf := tp.NewClientProfile()

	// 调试模式
	if os.Getenv("TDP_DEBUG") != "" {
		cpf.Debug = true
	}

	// 网络错误重试
	cpf.NetworkFailureMaxRetries = 2

	// API 限频重试
	cpf.RateLimitExceededMaxRetries = 2

	// 启用地域容灾
	cpf.DisableRegionBreaker = false
	cpf.BackupEndpoint = "ap-hongkong." + rp.RootDomain

	// 分地域接入，避免限频
	if rp.Endpoint != "" {
		cpf.HttpProfile.Endpoint = rp.Service + "." + rp.Endpoint + "." + rp.RootDomain
	} else if rp.Region != "" {
		cpf.HttpProfile.Endpoint = rp.Service + "." + rp.Region + "." + rp.RootDomain
	} else {
		cpf.HttpProfile.RootDomain = rp.RootDomain
	}

	cred := tc.NewCredential(rp.SecretId, rp.SecretKey)

	return tc.NewCommonClient(cred, rp.Region, cpf)

}
