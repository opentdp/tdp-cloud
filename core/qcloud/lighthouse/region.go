package lighthouse

import (
	"tdp-cloud/core/midware"

	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

// 查询地域列表

func DescribeRegions(ud *midware.Userdata) (*lighthouse.DescribeRegionsResponse, error) {

	client, _ := NewClient(ud)

	request := lighthouse.NewDescribeRegionsRequest()

	return client.DescribeRegions(request)

}
