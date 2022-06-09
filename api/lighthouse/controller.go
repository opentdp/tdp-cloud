package lighthouse

import (
	"tdp-cloud/core/midware"
	"tdp-cloud/core/qcloud/lighthouse"

	"github.com/gin-gonic/gin"
)

// 获取地域

func describeRegions(c *gin.Context) {

	ud := midware.GetUserdata(c)

	response, err := lighthouse.DescribeRegions(ud)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}

// 获取地域实例

func describeInstances(c *gin.Context) {

	ud := midware.GetUserdata(c)

	response, err := lighthouse.DescribeInstances(ud)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}

// 获取实例流量包

func describeInstancesTrafficPackages(c *gin.Context) {

	ud := midware.GetUserdata(c)

	response, err := lighthouse.DescribeInstancesTrafficPackages(ud)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}
