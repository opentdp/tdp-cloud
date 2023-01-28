package socket

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Socket(wsi *gin.RouterGroup) {

	rg := wsi.Group("/")

	// 匿名接口

	{
		rg.GET("/worker", worker)
	}

	// 需授权接口

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/ssh", ssh)
	}

}
