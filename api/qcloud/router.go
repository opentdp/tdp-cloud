package qcloud

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/qcloud")

	// 匿名接口

	{
		rg.GET("/vnc", vncProxy)
	}

	// 需授权接口

	rg.Use(midware.Auth())

	{
		rg.POST("/", apiProxy)
	}

}
