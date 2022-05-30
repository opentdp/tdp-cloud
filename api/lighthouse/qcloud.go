package lighthouse

import (
	"github.com/gin-gonic/gin"

	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"

	"tdp-cloud/core/qcloud"
)

// 查询地域列表

func describeRegions(c *gin.Context) {

	client := qcloud.NewLighthouseClient(c)

	request := lighthouse.NewDescribeRegionsRequest()
	response, err := client.DescribeRegions(request)

	c.Set("Payload", response.Response)
	c.Set("Error", err)

}

// 查看实例列表

func describeInstances(c *gin.Context) {

	client := qcloud.NewLighthouseClient(c)

	request := lighthouse.NewDescribeInstancesRequest()
	response, err := client.DescribeInstances(request)

	c.Set("Payload", response.Response)
	c.Set("Error", err)

}

//查看实例流量包详情

func describeInstancesTrafficPackages(c *gin.Context) {

	client := qcloud.NewLighthouseClient(c)

	request := lighthouse.NewDescribeInstancesTrafficPackagesRequest()
	response, err := client.DescribeInstancesTrafficPackages(request)

	c.Set("Payload", response.Response)
	c.Set("Error", err)

}
