package passport

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

var ctrl = &Controller{}

func Router(api *gin.RouterGroup) {

	rg := api.Group("/passport")

	// 匿名接口

	{
		rg.POST("/login", ctrl.login)
		rg.POST("/register", ctrl.register)
	}

	// 需授权接口

	rg.Use(midware.AuthGuard)

	{
		rg.POST("/profile", ctrl.profile)
		rg.POST("/profile/update", ctrl.profileUpdate)
		rg.POST("/avatar/update", ctrl.avatarUpdate)
		rg.POST("/summary", ctrl.summary)
	}

}
