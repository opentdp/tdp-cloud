package lighthouse

import (
	"tdp-cloud/core/qcloud/lighthouse"

	"github.com/gin-gonic/gin"
)

// 获取地域

func describeRegions(c *gin.Context) {

	config_, _ := c.Get("Config")
	config := config_.([3]string)

	response, err := lighthouse.DescribeRegions(config)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}

// 获取地域实例

func describeInstances(c *gin.Context) {

	config_, _ := c.Get("Config")
	config := config_.([3]string)

	response, err := lighthouse.DescribeInstances(config)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}

// 获取实例流量包

func describeInstancesTrafficPackages(c *gin.Context) {

	config_, _ := c.Get("Config")
	config := config_.([3]string)

	response, err := lighthouse.DescribeInstancesTrafficPackages(config)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}
