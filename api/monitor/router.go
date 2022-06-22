package monitor

import (
	"tdp-cloud/core/midware"

	"github.com/gin-gonic/gin"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/monitor")

	rg.Use(midware.Auth())
	rg.Use(midware.Secret())

	{
		rg.POST("/getMonitorData/:region", getMonitorData)
	}

}
