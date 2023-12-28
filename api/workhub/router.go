package workhub

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	ctrl := &Controller{}

	rg := api.Group("/workhub")

	// 需授权接口

	rg.Use(midware.AuthGuard)

	{
		rg.POST("/list", ctrl.list)
	}

	// 管理员接口

	rg.Use(midware.AdminGuard)

	{
		rg.POST("/detail", ctrl.detail)
		rg.POST("/exec", ctrl.exec)
		rg.POST("/filer", ctrl.filer)
	}

}

func Socket(wsi *gin.RouterGroup) {

	ctrl := &Controller{}

	rg := wsi.Group("/")

	{
		rg.GET("/workhub", ctrl.join)
		rg.GET("/workhub/:mid", ctrl.join)
	}

}
