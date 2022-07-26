package lighthouse

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
	"tdp-cloud/core/qcloud/lighthouse"
)

// 获取地域

func describeRegions(c *gin.Context) {

	var ud = midware.GetUserdata(c)

	if res, err := lighthouse.DescribeRegions(ud); err == nil {
		c.Set("Payload", res.Response)
	} else {
		c.Set("Error", err)
	}

}

// 获取地域实例

func describeInstances(c *gin.Context) {

	var ud = midware.GetUserdata(c)
	var rq lighthouse.DescribeInstancesRequestParams

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	if res, err := lighthouse.DescribeInstances(ud, &rq); err == nil {
		c.Set("Payload", res.Response)
	} else {
		c.Set("Error", err)
	}

}

// 获取实例流量包

func describeInstancesTrafficPackages(c *gin.Context) {

	var ud = midware.GetUserdata(c)
	var rq lighthouse.DescribeInstancesTrafficPackagesRequestParams

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	if res, err := lighthouse.DescribeInstancesTrafficPackages(ud, &rq); err == nil {
		c.Set("Payload", res.Response)
	} else {
		c.Set("Error", err)
	}

}

// 获取实例流量包

func describeSnapshots(c *gin.Context) {

	var ud = midware.GetUserdata(c)
	var rq lighthouse.DescribeSnapshotsRequestParams

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	if res, err := lighthouse.DescribeSnapshots(ud, &rq); err == nil {
		c.Set("Payload", res.Response)
	} else {
		c.Set("Error", err)
	}

}

// 查询防火墙规则

func describeFirewallRules(c *gin.Context) {

	var ud = midware.GetUserdata(c)
	var rq lighthouse.DescribeFirewallRulesRequestParams

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	if res, err := lighthouse.DescribeFirewallRules(ud, &rq); err == nil {
		c.Set("Payload", res.Response)
	} else {
		c.Set("Error", err)
	}

}
