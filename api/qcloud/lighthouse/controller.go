package lighthouse

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
	"tdp-cloud/core/qcloud/lighthouse"
)

// 查询地域列表

func describeRegions(c *gin.Context) {

	var ud = midware.GetUserdata(c)

	if res, err := lighthouse.DescribeRegions(ud); err == nil {
		c.Set("Payload", res.Response)
	} else {
		c.Set("Error", err)
	}

}

// 查看实例列表

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

// 查询实例管理终端地址

func describeInstanceVncUrl(c *gin.Context) {

	var ud = midware.GetUserdata(c)
	var rq lighthouse.DescribeInstanceVncUrlRequestParams

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	if res, err := lighthouse.DescribeInstanceVncUrl(ud, &rq); err == nil {
		c.Set("Payload", res.Response)
	} else {
		c.Set("Error", err)
	}

}

// 查看实例流量包详情

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

// 查看快照列表

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
