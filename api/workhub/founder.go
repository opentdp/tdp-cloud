package workhub

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/helper/psutil"
)

// 主机信息

func host(c *gin.Context) {

	info := psutil.Detail(false)

	c.Set("Payload", gin.H{"Stat": info})

}

// 主机IP

func hostIp(c *gin.Context) {

	ipv4, ipv6 := psutil.PublicAddress(false)

	c.Set("Payload", gin.H{"Ipv4": ipv4, "Ipv6": ipv6})

}
