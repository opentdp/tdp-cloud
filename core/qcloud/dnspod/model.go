package dnspod

import (
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"

	"tdp-cloud/core/qcloud"
)

// 创建客户端

func NewClient(config [3]string) (*dnspod.Client, error) {

	credential, cpf := qcloud.NewCredentialProfile(config)

	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"

	client, err := dnspod.NewClient(credential, "", cpf)

	return client, err

}

// 获取域名列表

func DescribeDomainList(config [3]string) (*dnspod.DescribeDomainListResponse, error) {

	client, err := NewClient(config)

	if err != nil {
		return nil, err
	}

	request := dnspod.NewDescribeDomainListRequest()
	response, _ := client.DescribeDomainList(request)

	return response, err

}
