package lighthouse

import (
	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"

	"tdp-cloud/core/midware"
)

// 查询地域列表

func DescribeRegions(ud *midware.Userdata) (*lighthouse.DescribeRegionsResponse, error) {

	client, _ := NewClient(ud)

	request := lighthouse.NewDescribeRegionsRequest()

	return client.DescribeRegions(request)

}
