package workhub

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/workhub")

	// 需授权接口

	rg.Use(midware.AuthGuard)

	{
		rg.POST("/list", list)
		rg.POST("/detail/:id", detail)
		rg.POST("/exec/:id", exec)
	}

	// 管理员接口

	rg.Use(midware.AdminGuard)

	{
		rg.POST("/host", host)
		rg.POST("/host/ip", hostIp)
	}

}

func Socket(wsi *gin.RouterGroup) {

	rg := wsi.Group("/")

	{
		rg.GET("/workhub", register)
		rg.GET("/workhub/:mid", register)
	}

}
