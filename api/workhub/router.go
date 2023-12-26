package workhub

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

var (
	ctrl = &Controller{}
	host = &HostController{}
	node = &NodeController{}
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/workhub")

	// 需授权接口

	rg.Use(midware.AuthGuard)

	{
		rg.POST("/list", ctrl.list)
		rg.POST("/detail/:id", node.detail)
		rg.POST("/exec/:id", node.exec)
		rg.POST("/filer/:id", node.filer)
	}

	// 管理员接口

	rg.Use(midware.AdminGuard)

	{
		rg.POST("/detail", host.detail)
		rg.POST("/exec", host.exec)
		rg.POST("/filer", host.filer)
	}

}

func Socket(wsi *gin.RouterGroup) {

	rg := wsi.Group("/")

	{
		rg.GET("/workhub", ctrl.join)
		rg.GET("/workhub/:mid", ctrl.join)
	}

}
