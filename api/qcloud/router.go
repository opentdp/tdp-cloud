package qcloud

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"

	"tdp-cloud/api/qcloud/cam"
	"tdp-cloud/api/qcloud/dnspod"
	"tdp-cloud/api/qcloud/lighthouse"
	"tdp-cloud/api/qcloud/monitor"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/qcloud")

	rg.Use(midware.Auth())
	rg.Use(midware.Secret())

	{
		cam.Router(rg)
		dnspod.Router(rg)
		lighthouse.Router(rg)
		monitor.Router(rg)
	}

}
