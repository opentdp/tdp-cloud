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
		c.Set("Error", err)
		return
	}

	response, err := dnspod.DescribeDomainList(ud, &rq)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}

// 获取解析记录列表

func describeRecordList(c *gin.Context) {

	var ud = midware.GetUserdata(c)
	var rq dnspod.DescribeRecordListRequestParams

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	response, err := dnspod.DescribeRecordList(ud, &rq)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}

// 获取等级允许的线路

func describeRecordLineList(c *gin.Context) {

	var ud = midware.GetUserdata(c)
	var rq dnspod.DescribeRecordLineListRequestParams

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	response, err := dnspod.DescribeRecordLineList(ud, &rq)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}

// 获取等级允许的记录类型

func describeRecordType(c *gin.Context) {

	var ud = midware.GetUserdata(c)
	var rq dnspod.DescribeRecordTypeRequestParams

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	response, err := dnspod.DescribeRecordType(ud, &rq)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}

// 修改解析记录

func modifyRecord(c *gin.Context) {

	var ud = midware.GetUserdata(c)
	var rq dnspod.ModifyRecordRequestParams

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	response, err := dnspod.ModifyRecord(ud, &rq)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}
