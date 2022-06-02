package lighthouse

import (
	"github.com/gin-gonic/gin"

	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"

	"tdp-cloud/core/qcloud"
)

// 获取所有地域和实例列表

func getAllRegionsInstances(c *gin.Context) {

	var result = make(map[string]interface{})

	// 获取所有地域
	regionsClient := qcloud.NewLighthouseClient(c, "")
	regionsRequest := lighthouse.NewDescribeRegionsRequest()
	regionsResponse, err := regionsClient.DescribeRegions(regionsRequest)

	// 获取所有地域的实例
	if err == nil && regionsResponse.Response.RegionSet != nil {
		var instanceSet []*lighthouse.Instance

		for _, region := range regionsResponse.Response.RegionSet {
			regionsClient := qcloud.NewLighthouseClient(c, *region.Region)
			instancesRequest := lighthouse.NewDescribeInstancesRequest()
			instanceResponse, er2 := regionsClient.DescribeInstances(instancesRequest)
			if er2 == nil && instanceResponse.Response.InstanceSet != nil {
				instanceSet = append(instanceSet, instanceResponse.Response.InstanceSet...)
			}
		}

		result["RegionSet"] = regionsResponse.Response.RegionSet
		result["instanceSet"] = instanceSet
	}

	c.Set("Payload", result)
	c.Set("Error", err)

}
