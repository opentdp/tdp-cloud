package workhub

import (
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/opentdp/go-helper/psutil"
)

// 主机信息

func hostDetail(c *gin.Context) {

	c.Set("Payload", gin.H{
		"Stat":         psutil.Detail(true),
		"MemStat":      psutil.GoMemory(),
		"NumGoroutine": runtime.NumGoroutine(),
	})

}
