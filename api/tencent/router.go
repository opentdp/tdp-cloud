package tencent

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

var ctrl = &Controller{}

func Router(api *gin.RouterGroup) {

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
