package user

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	ctrl := &Controller{}

	rg := api.Group("/user")

	// 需授权接口

	rg.Use(midware.AuthGuard)

	// 管理员接口

	rg.Use(midware.AdminGuard)

	{
		rg.POST("/list", ctrl.list)
		rg.POST("/create", ctrl.create)
		rg.POST("/detail", ctrl.detail)
		rg.POST("/update", ctrl.update)
		rg.POST("/delete", ctrl.delete)
	}

}
