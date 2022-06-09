package lighthouse

import (
	"tdp-cloud/core/midware"

	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

// 查看实例列表

type DescribeInstancesRequest struct {
	InstanceIds []*string
	Filters     []*lighthouse.Filter
	Offset      *int64
	Limit       *int64
}

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

type DescribeInstancesTrafficPackagesRequest struct {
	InstanceIds []*string
	Offset      *int64
	Limit       *int64
}

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
