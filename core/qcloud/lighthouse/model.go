package lighthouse

import (
	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"

	"tdp-cloud/core/qcloud"
)

// 创建客户端

func NewClient(config [3]string) (*lighthouse.Client, error) {

	credential, cpf := qcloud.NewCredentialProfile(config)

	if config[2] == "" {
		cpf.HttpProfile.Endpoint = "lighthouse.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "lighthouse." + config[2] + ".tencentcloudapi.com"
	}

	return lighthouse.NewClient(credential, config[2], cpf)

}

// 查询地域列表

func DescribeRegions(config [3]string) (*lighthouse.DescribeRegionsResponse, error) {

	client, _ := NewClient(config)

	request := lighthouse.NewDescribeRegionsRequest()

	return client.DescribeRegions(request)

}

// 查看实例列表

func DescribeInstances(config [3]string) (*lighthouse.DescribeInstancesResponse, error) {

	client, _ := NewClient(config)

	request := lighthouse.NewDescribeInstancesRequest()

	return client.DescribeInstances(request)

}

// 查看实例流量包详情

func DescribeInstancesTrafficPackages(config [3]string) (*lighthouse.DescribeInstancesTrafficPackagesResponse, error) {

	client, _ := NewClient(config)

	request := lighthouse.NewDescribeInstancesTrafficPackagesRequest()

	return client.DescribeInstancesTrafficPackages(request)

}
