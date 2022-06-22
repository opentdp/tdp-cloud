package lighthouse

import (
	"tdp-cloud/core/midware"
	"tdp-cloud/core/qcloud/lighthouse"

	"github.com/gin-gonic/gin"
)

// 获取地域

func describeRegions(c *gin.Context) {

	var ud = midware.GetUserdata(c)

	response, err := lighthouse.DescribeRegions(ud)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}

// 获取地域实例

func describeInstances(c *gin.Context) {

	var ud = midware.GetUserdata(c)
	var rq lighthouse.DescribeInstancesRequestParams

	c.ShouldBind(&rq)

	response, err := lighthouse.DescribeInstances(ud, &rq)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}

// 获取实例流量包

func describeInstancesTrafficPackages(c *gin.Context) {

	var ud = midware.GetUserdata(c)
	var rq lighthouse.DescribeInstancesTrafficPackagesRequestParams

	c.ShouldBind(&rq)

	response, err := lighthouse.DescribeInstancesTrafficPackages(ud, &rq)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}

// 获取实例流量包

func describeSnapshots(c *gin.Context) {

	var ud = midware.GetUserdata(c)
	var rq lighthouse.DescribeSnapshotsRequestParams

	c.ShouldBind(&rq)

	response, err := lighthouse.DescribeSnapshots(ud, &rq)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}

// 查询防火墙规则

func describeFirewallRules(c *gin.Context) {

	var ud = midware.GetUserdata(c)
	var rq lighthouse.DescribeFirewallRulesRequestParams

	c.ShouldBind(&rq)

	response, err := lighthouse.DescribeFirewallRules(ud, &rq)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}
