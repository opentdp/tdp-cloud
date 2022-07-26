package monitor

import (
	"github.com/gin-gonic/gin"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/monitor")

	{
		rg.POST("/getMonitorData/:region", getMonitorData)
	}

}
