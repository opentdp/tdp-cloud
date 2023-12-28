package tencent

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	ctrl := &Controller{}

	rg := api.Group("/tencent")

	// 匿名接口

	{
		rg.GET("/vnc", ctrl.vncProxy)
	}

	// 需授权接口

	rg.Use(midware.AuthGuard)

	{
		rg.POST("/:id", ctrl.apiProxy)
	}

}
