package lighthouse

import (
	"github.com/gin-gonic/gin"

	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

// 获取所有地域和实例列表

func getAllRegionsInstances(c *gin.Context) {

	regionSet := DescribeRegions(c)

	instanceSet := DescribeInstances(c, regionSet)

	var result = make(map[string]interface{})
	result["RegionSet"] = regionSet
	result["InstanceSet"] = instanceSet

	c.Set("Payload", result)

}

//查看实例流量包详情

func describeInstancesTrafficPackages(c *gin.Context) {

	region := c.Param("region")
	client := NewClient(c, region)

	request := lighthouse.NewDescribeInstancesTrafficPackagesRequest()
	response, err := client.DescribeInstancesTrafficPackages(request)

	c.Set("Payload", response.Response)
	c.Set("Error", err)

}
