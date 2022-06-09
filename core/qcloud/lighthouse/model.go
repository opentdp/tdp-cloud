package lighthouse

import (
	"tdp-cloud/core/midware"
	"tdp-cloud/core/qcloud"

	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

// 创建客户端

func NewClient(ud midware.Userdata) (*lighthouse.Client, error) {

	credential, cpf := qcloud.NewCredentialProfile(ud)

	if ud.Region == "" {
		cpf.HttpProfile.Endpoint = "lighthouse.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "lighthouse." + ud.Region + ".tencentcloudapi.com"
	}

	return lighthouse.NewClient(credential, ud.Region, cpf)

}

// 查询地域列表

func DescribeRegions(ud midware.Userdata) (*lighthouse.DescribeRegionsResponse, error) {

	client, _ := NewClient(ud)

	request := lighthouse.NewDescribeRegionsRequest()

	return client.DescribeRegions(request)

}

// 查看实例列表

func DescribeInstances(ud midware.Userdata) (*lighthouse.DescribeInstancesResponse, error) {

	client, _ := NewClient(ud)

	request := lighthouse.NewDescribeInstancesRequest()

	return client.DescribeInstances(request)

}

// 查看实例流量包详情

func DescribeInstancesTrafficPackages(ud midware.Userdata) (*lighthouse.DescribeInstancesTrafficPackagesResponse, error) {

	client, _ := NewClient(ud)

	request := lighthouse.NewDescribeInstancesTrafficPackagesRequest()

	return client.DescribeInstancesTrafficPackages(request)

}
