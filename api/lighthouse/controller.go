package lighthouse

import (
	"tdp-cloud/core/qcloud/lighthouse"

	"github.com/gin-gonic/gin"
)

// 获取地域列表

func describeRegions(c *gin.Context) {

	config_, _ := c.Get("Config")
	config := config_.([3]string)

	response, err := lighthouse.DescribeRegions(config)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}

// 获取地域实例列表

func DescribeInstances(c *gin.Context) {

	config_, _ := c.Get("Config")
	config := config_.([3]string)

	response, err := lighthouse.DescribeInstances(config)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}

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

// 获取实例流量包

func describeTrafficPackages(c *gin.Context) {

	config_, _ := c.Get("Config")
	config := config_.([3]string)

	response, err := lighthouse.DescribeTrafficPackages(config)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}

// 获取所有地域实例流量包

func describeRegionsTrafficPackages(c *gin.Context) {

	config_, _ := c.Get("Config")
	config := config_.([3]string)

	regionResponse, err := lighthouse.DescribeRegions(config)

	if err != nil {
		c.Set("Error", err)
		return
	}

	regionSet := regionResponse.Response.RegionSet

	trafficPackageSet, ers := lighthouse.DescribeTrafficPackagesAll(config, regionSet)
	response := gin.H{"RegionSet": regionSet, "trafficPackageSet": trafficPackageSet}

	c.Set("Payload", response)

	if len(ers) > 0 {
		c.Set("Error", ers)
	}

}
