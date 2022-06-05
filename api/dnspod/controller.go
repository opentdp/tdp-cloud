package dnspod

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/qcloud/dnspod"
)

// 获取域名列表

func describeDomainList(c *gin.Context) {

	config_, _ := c.Get("Config")
	config := config_.([3]string)

	response, err := dnspod.DescribeDomainList(config)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}
