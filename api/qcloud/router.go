package qcloud

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/qcloud")

	rg.Use(midware.Auth()).POST("/", doRequest)

}
