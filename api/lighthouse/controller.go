package lighthouse

import (
	"tdp-cloud/core/qcloud/lighthouse"

	"github.com/gin-gonic/gin"
)

// 获取所有地域和实例列表

func describeRegionsInstances(c *gin.Context) {

	config_, _ := c.Get("Config")
	config := config_.([3]string)

	regionResponse, err := lighthouse.DescribeRegions(config)

	if err != nil {
		c.Set("Error", err)
		return
	}

	regionSet := regionResponse.Response.RegionSet

	instanceSet, ers := lighthouse.DescribeInstancesAll(config, regionSet)
	response := gin.H{"RegionSet": regionSet, "InstanceSet": instanceSet}

	c.Set("Payload", response)

	if len(ers) > 0 {
		c.Set("Error", ers)
	}

}

// 获取实例流量包详情

func describeInstancesTrafficPackages(c *gin.Context) {

	config_, _ := c.Get("Config")
	config := config_.([3]string)

	response, err := lighthouse.DescribeInstancesTrafficPackages(config)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}

// 获取所有地域实例流量包详情

func DescribeInstancesTrafficPackagesAll(c *gin.Context) {

	config_, _ := c.Get("Config")
	config := config_.([3]string)

	regionResponse, err := lighthouse.DescribeRegions(config)

	if err != nil {
		c.Set("Error", err)
		return
	}

	regionSet := regionResponse.Response.RegionSet

	instanceSet, ers := lighthouse.DescribeInstancesTrafficPackagesAll(config, regionSet)
	response := gin.H{"RegionSet": regionSet, "InstanceSet": instanceSet}

	c.Set("Payload", response)

	if len(ers) > 0 {
		c.Set("Error", ers)
	}

}
