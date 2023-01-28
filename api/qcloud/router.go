package qcloud

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/")

	// 匿名接口

	{
		rg.GET("/qcloud/vnc", vncProxy)
	}

	// 需授权接口

	rg.Use(midware.AuthGuard())

	{
		rg.POST("/qcloud", apiProxy)
	}

}
