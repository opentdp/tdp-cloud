package tencent

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"

	"github.com/mitchellh/mapstructure"

	tc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	te "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	th "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	tp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"

	"tdp-cloud/cmd/args"
)

func Request(rq *ReqeustParam) (any, error) {

	resp, err := newClient(rq)

	if err != nil {
		return nil, getSDKError(err)
	}

	body := resp.GetBody()
	res := &ResponseData{}

	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res.Response, nil

}

func newClient(rq *ReqeustParam) (*th.CommonResponse, error) {

	cpf := tp.NewClientProfile()

	// 调试开关
	cpf.Debug = args.Debug

	// 网络错误重试
	cpf.NetworkFailureMaxRetries = 2

	// API 限频重试
	cpf.RateLimitExceededMaxRetries = 2

	// 启用地域容灾
	cpf.DisableRegionBreaker = false
	cpf.BackupEndpoint = "ap-hongkong." + th.RootDomain

	// 按地域设置接口
	if rq.Endpoint != "" {
		cpf.HttpProfile.Endpoint = rq.Endpoint // 完整域名
	} else if rq.Region != "" && !strings.HasSuffix(rq.Region, "-ec") {
		cpf.HttpProfile.Endpoint = rq.Service + "." + rq.Region + "." + th.RootDomain
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
