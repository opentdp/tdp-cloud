package user

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/user")

	// 匿名接口

	{
		rg.POST("/login", login)
		rg.POST("/register", register)
	}

	// 需授权接口

	rg.Use(midware.Auth())

	{
		rg.PATCH("/info", updateInfo)
		rg.PATCH("/password", updatePassword)
	}

}
