package monitor

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/monitor")

	rg.Use(midware.Auth())
	rg.Use(midware.Secret())

	{
		rg.POST("/getMonitorData/:region", getMonitorData)
	}

}
