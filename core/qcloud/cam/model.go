package cam

import (
	"tdp-cloud/core/midware"
	"tdp-cloud/core/qcloud"

	cam "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cam/v20190116"
)

// 创建客户端

func NewClient(ud midware.Userdata) (*cam.Client, error) {

	credential, cpf := qcloud.NewCredentialProfile(ud)

	return cam.NewClient(credential, "", cpf)

}

// 查询账户摘要

func GetAccountSummary(ud midware.Userdata) (*cam.GetAccountSummaryResponse, error) {

	client, _ := NewClient(ud)

	request := cam.NewGetAccountSummaryRequest()

	return client.GetAccountSummary(request)

}
