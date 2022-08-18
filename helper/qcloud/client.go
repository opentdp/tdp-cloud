package qcloud

import (
	"encoding/json"
	"os"

	tc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	th "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	tp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

type Params struct {
	Service    string `note:"产品名称"`
	Version    string `note:"接口版本"`
	Action     string `note:"接口名称"`
	Payload    []byte `note:"结构化数据"`
	Region     string `note:"资源所在区域"`
	Endpoint   string `note:"指定接口区域"`
	SecretId   string `note:"访问密钥 Id"`
	SecretKey  string `note:"访问密钥 Key"`
	RootDomain string `note:"API 根域名"`
}

type Response struct {
	Response any
}

func NewRequest(rp *Params) (*Response, error) {

	request := th.NewCommonRequest(rp.Service, rp.Version, rp.Action)

	if rp.Payload != nil {
		request.SetActionParameters(rp.Payload)
	}

	client := NewClient(rp)
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

func NewClient(rp *Params) *tc.Client {

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
