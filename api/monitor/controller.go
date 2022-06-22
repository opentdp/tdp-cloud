package monitor

import (
	"tdp-cloud/core/midware"
	"tdp-cloud/core/qcloud/monitor"

	"github.com/gin-gonic/gin"
)

// 获取地域

func getMonitorData(c *gin.Context) {

	var ud = midware.GetUserdata(c)
	var rq monitor.GetMonitorDataRequestParams

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", "参数错误")
		return
	}

	response, err := monitor.GetMonitorData(ud, &rq)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}
