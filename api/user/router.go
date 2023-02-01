package user

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/user")

	// 匿名接口

	{
		rg.POST("/login", login)
		rg.POST("/register", create)
	}

	// 需授权接口

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/info", detail)
		rg.PATCH("/info", update)
		rg.PATCH("/password", updatePassword)
	}

}
