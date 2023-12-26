package upgrade

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

var ctrl = &Controller{}

func Router(api *gin.RouterGroup) {

	rg := api.Group("/upgrade")

	// 需授权接口

	rg.Use(midware.AuthGuard)

	// 管理员接口

	rg.Use(midware.AdminGuard)

	{
		rg.POST("/check", ctrl.check)
		rg.POST("/apply", ctrl.apply)
	}

}
