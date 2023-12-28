package upgrade

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	ctrl := &Controller{}

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
