package dnspod

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
	"tdp-cloud/core/qcloud/dnspod"
)

// 获取域名列表

func describeDomainList(c *gin.Context) {

	ud := midware.GetUserdata(c)

	response, err := dnspod.DescribeDomainList(ud)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}
