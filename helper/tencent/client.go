package tencent

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"

	tc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	te "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	th "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	tp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func Request(rq *Params) (any, error) {

	resp, err := newClient(rq)

	if err != nil {
		return nil, getSDKError(err)
	}

	body := resp.GetBody()
	res := &Result{}

	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res.Response, nil

}

func newClient(rq *Params) (*th.CommonResponse, error) {

	cpf := tp.NewClientProfile()

	// 调试开关
	cpf.Debug = viper.GetBool("debug")

	// 接口根域名
	if rq.RootDomain == "" {
		rq.RootDomain = "tencentcloudapi.com"
	}

	// 网络错误重试
	cpf.NetworkFailureMaxRetries = 2

	// API 限频重试
	cpf.RateLimitExceededMaxRetries = 2

	// 启用地域容灾
	cpf.DisableRegionBreaker = false
	cpf.BackupEndpoint = "ap-hongkong." + rq.RootDomain

	// 按地域设置接口
	if rq.Endpoint != "" {
		cpf.HttpProfile.Endpoint = rq.Service + "." + rq.Endpoint + "." + rq.RootDomain
	} else if rq.Region != "" {
		cpf.HttpProfile.Endpoint = rq.Service + "." + rq.Region + "." + rq.RootDomain
	} else {
		cpf.HttpProfile.RootDomain = rq.RootDomain
	}

	// 初始化客户端
	cred := tc.NewCredential(rq.SecretId, rq.SecretKey)
	client := tc.NewCommonClient(cred, rq.Region, cpf)

	// 构造请求信息
	request := th.NewCommonRequest(rq.Service, rq.Version, rq.Action)
	if rq.Payload != nil {
		request.SetActionParameters(rq.Payload)
	}

	// 发起请求
	response := th.NewCommonResponse()
	err := client.Send(request, response)

	return response, err

}

func getSDKError(e error) error {

	se := te.TencentCloudSDKError{}

	if mapstructure.Decode(e, &se) != nil {
		return e
	}

	if se.Message == "" {
		return e
	}

	exp := regexp.MustCompile(`\[request id:.+\]`)
	msg := strings.Split(exp.ReplaceAllString(se.Message, ""), "\n")[0]

	return errors.New(msg)

}
