package lighthouse

import (
	"tdp-cloud/core/midware"
	"tdp-cloud/core/qcloud"

	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

// 创建客户端

func NewClient(ud *midware.Userdata) (*lighthouse.Client, error) {

	credential, cpf := qcloud.NewCredentialProfile(ud)

	if ud.Region != "" {
		cpf.HttpProfile.Endpoint = "lighthouse." + ud.Region + ".tencentcloudapi.com"
	}

	return lighthouse.NewClient(credential, ud.Region, cpf)

}

// 查询地域列表

func DescribeRegions(ud *midware.Userdata) (*lighthouse.DescribeRegionsResponse, error) {

	client, _ := NewClient(ud)

	request := lighthouse.NewDescribeRegionsRequest()

	return client.DescribeRegions(request)

}

// 查看实例列表

func DescribeInstances(ud *midware.Userdata, rq *DescribeInstancesRequest) (*lighthouse.DescribeInstancesResponse, error) {

	client, _ := NewClient(ud)

	request := lighthouse.NewDescribeInstancesRequest()

	if len(rq.InstanceIds) > 0 {
		request.InstanceIds = rq.InstanceIds
	}

	if len(rq.Filters) > 0 {
		request.Filters = rq.Filters
	}

	if rq.Offset != nil {
		request.Offset = rq.Offset
	}

	if rq.Limit != nil {
		request.Limit = rq.Limit
	}

	return client.DescribeInstances(request)

}

// 查看实例流量包详情

func DescribeInstancesTrafficPackages(ud *midware.Userdata, rq *DescribeInstancesTrafficPackagesRequest) (*lighthouse.DescribeInstancesTrafficPackagesResponse, error) {

	client, _ := NewClient(ud)

	request := lighthouse.NewDescribeInstancesTrafficPackagesRequest()

	if len(rq.InstanceIds) > 0 {
		request.InstanceIds = rq.InstanceIds
	}

	if rq.Offset != nil {
		request.Offset = rq.Offset
	}

	if rq.Limit != nil {
		request.Limit = rq.Limit
	}

	return client.DescribeInstancesTrafficPackages(request)

}
