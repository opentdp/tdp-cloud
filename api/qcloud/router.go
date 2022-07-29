package qcloud

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/qcloud")

	rg.Use(midware.Auth())
	rg.Use(midware.Secret())

	rg.POST("/:service/:version/:action", doRequest)
	rg.POST("/:service/:version/:action/:region", doRequest)

}
