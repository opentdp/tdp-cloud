package workhub

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/helper/psutil"
)

// 主机信息

func host(c *gin.Context) {

	info := psutil.Detail()

	c.Set("Payload", gin.H{"Stat": info})

}

// 主机IP

func hostIp(c *gin.Context) {

	ip := psutil.PublicIpAddress(false)

	c.Set("Payload", gin.H{"Ip": ip})

}
