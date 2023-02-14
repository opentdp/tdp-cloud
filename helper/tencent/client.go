package tencent

import (
	"encoding/json"
	"errors"
	"os"
	"regexp"
	"strings"

	"github.com/mitchellh/mapstructure"

	tc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	te "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	th "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	tp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func Request(rp *Params) (any, error) {

	resp, err := newClient(rp)

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

func newClient(rp *Params) (*th.CommonResponse, error) {

	cpf := tp.NewClientProfile()

	// 调试模式
	if os.Getenv("TDP_DEBUG") != "" {
		cpf.Debug = true
	}

	// 接口根域名
	if rp.RootDomain == "" {
		rp.RootDomain = "tencentcloudapi.com"
	}

	// 网络错误重试
	cpf.NetworkFailureMaxRetries = 2

	// API 限频重试
	cpf.RateLimitExceededMaxRetries = 2

	// 启用地域容灾
	cpf.DisableRegionBreaker = false
	cpf.BackupEndpoint = "ap-hongkong." + rp.RootDomain

	// 按地域设置接口
	if rp.Endpoint != "" {
		cpf.HttpProfile.Endpoint = rp.Service + "." + rp.Endpoint + "." + rp.RootDomain
	} else if rp.Region != "" {
		cpf.HttpProfile.Endpoint = rp.Service + "." + rp.Region + "." + rp.RootDomain
	} else {
		cpf.HttpProfile.RootDomain = rp.RootDomain
	}

	// 初始化客户端
	cred := tc.NewCredential(rp.SecretId, rp.SecretKey)
	client := tc.NewCommonClient(cred, rp.Region, cpf)

	// 构造请求信息
	request := th.NewCommonRequest(rp.Service, rp.Version, rp.Action)
	if rp.Payload != nil {
		request.SetActionParameters(rp.Payload)
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

	re, _ := regexp.Compile(`\[request id:.+\]`)
	msg := strings.Split(re.ReplaceAllString(se.Message, ""), "\n")[0]

	return errors.New(msg)

}
