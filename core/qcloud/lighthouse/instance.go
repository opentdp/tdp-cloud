package lighthouse

import (
	"tdp-cloud/core/midware"

	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

// 查看实例列表

type DescribeInstancesRequestParams = lighthouse.DescribeInstancesRequestParams

func DescribeInstances(ud *midware.Userdata, rq *DescribeInstancesRequestParams) (*lighthouse.DescribeInstancesResponse, error) {

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

type DescribeInstancesTrafficPackagesRequestParams = lighthouse.DescribeInstancesTrafficPackagesRequestParams

func DescribeInstancesTrafficPackages(ud *midware.Userdata, rq *DescribeInstancesTrafficPackagesRequestParams) (*lighthouse.DescribeInstancesTrafficPackagesResponse, error) {

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
