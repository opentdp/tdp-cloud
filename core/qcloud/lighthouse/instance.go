package lighthouse

import (
	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"

	"tdp-cloud/core/midware"
	"tdp-cloud/core/utils"
)

// 查看实例列表

type DescribeInstancesRequestParams = lighthouse.DescribeInstancesRequestParams

func DescribeInstances(ud *midware.Userdata, rq *DescribeInstancesRequestParams) (*lighthouse.DescribeInstancesResponse, error) {

	client, _ := NewClient(ud)

	request := lighthouse.NewDescribeInstancesRequest()
	request.FromJsonString(utils.ToJsonString(rq))

	return client.DescribeInstances(request)

}

// 查看实例流量包详情

type DescribeInstancesTrafficPackagesRequestParams = lighthouse.DescribeInstancesTrafficPackagesRequestParams

func DescribeInstancesTrafficPackages(ud *midware.Userdata, rq *DescribeInstancesTrafficPackagesRequestParams) (*lighthouse.DescribeInstancesTrafficPackagesResponse, error) {

	client, _ := NewClient(ud)

	request := lighthouse.NewDescribeInstancesTrafficPackagesRequest()
	request.FromJsonString(utils.ToJsonString(rq))

	return client.DescribeInstancesTrafficPackages(request)

}
