package cam

import (
	cam "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cam/v20190116"

	"tdp-cloud/core/qcloud"
)

// 创建客户端

func NewClient(config [3]string) (*cam.Client, error) {

	credential, cpf := qcloud.NewCredentialProfile(config)

	cpf.HttpProfile.Endpoint = "cam.tencentcloudapi.com"

	client, err := cam.NewClient(credential, "", cpf)

	return client, err

}

// 获取账号概要信息

func GetAccountSummary(config [3]string) (*cam.GetAccountSummaryResponse, error) {

	client, err := NewClient(config)

	if err != nil {
		return nil, err
	}

	request := cam.NewGetAccountSummaryRequest()
	response, err := client.GetAccountSummary(request)

	return response, err

}
