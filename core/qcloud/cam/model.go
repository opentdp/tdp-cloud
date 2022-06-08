package cam

import (
	cam "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cam/v20190116"

	"tdp-cloud/core/qcloud"
)

// 创建客户端

func NewClient(config [3]string) (*cam.Client, error) {

	credential, cpf := qcloud.NewCredentialProfile(config)

	cpf.HttpProfile.Endpoint = "cam.tencentcloudapi.com"

	return cam.NewClient(credential, "", cpf)

}

// 查询账户摘要

func GetAccountSummary(config [3]string) (*cam.GetAccountSummaryResponse, error) {

	client, _ := NewClient(config)

	request := cam.NewGetAccountSummaryRequest()

	return client.GetAccountSummary(request)

}
