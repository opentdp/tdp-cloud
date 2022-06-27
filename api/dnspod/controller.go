package dnspod

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
	"tdp-cloud/core/qcloud/dnspod"
)

// 获取域名列表

func describeDomainList(c *gin.Context) {

	var ud = midware.GetUserdata(c)
	var rq dnspod.DescribeDomainListRequestParams

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	if res, err := dnspod.DescribeDomainList(ud, &rq); err == nil {
		c.Set("Payload", res.Response)
	} else {
		c.Set("Error", err)
	}

}

// 获取解析记录列表

func describeRecordList(c *gin.Context) {

	var ud = midware.GetUserdata(c)
	var rq dnspod.DescribeRecordListRequestParams

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	if res, err := dnspod.DescribeRecordList(ud, &rq); err == nil {
		c.Set("Payload", res.Response)
	} else {
		c.Set("Error", err)
	}

}

// 获取等级允许的线路

func describeRecordLineList(c *gin.Context) {

	var ud = midware.GetUserdata(c)
	var rq dnspod.DescribeRecordLineListRequestParams

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	if res, err := dnspod.DescribeRecordLineList(ud, &rq); err == nil {
		c.Set("Payload", res.Response)
	} else {
		c.Set("Error", err)
	}

}

// 获取等级允许的记录类型

func describeRecordType(c *gin.Context) {

	var ud = midware.GetUserdata(c)
	var rq dnspod.DescribeRecordTypeRequestParams

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	if res, err := dnspod.DescribeRecordType(ud, &rq); err == nil {
		c.Set("Payload", res.Response)
	} else {
		c.Set("Error", err)
	}

}

// 修改解析记录

func modifyRecord(c *gin.Context) {

	var ud = midware.GetUserdata(c)
	var rq dnspod.ModifyRecordRequestParams

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	if res, err := dnspod.ModifyRecord(ud, &rq); err == nil {
		c.Set("Payload", res.Response)
	} else {
		c.Set("Error", err)
	}

}
