package workhub

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/helper/psutil"
)

// 主机信息

func host(c *gin.Context) {

	var rq *HostInfoParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	stat := psutil.Detail(rq.WithAddr)

	c.Set("Payload", gin.H{"Stat": stat})

}

// 主机IP

func hostIp(c *gin.Context) {

	var rq *HostInfoParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	ipv4, ipv6 := psutil.PublicAddress(rq.Force)

	c.Set("Payload", gin.H{"Ipv4": ipv4, "Ipv6": ipv6})

}

// 请求参数

type HostInfoParam struct {
	Force    bool
	WithAddr bool
}
