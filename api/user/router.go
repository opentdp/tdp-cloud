package user

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	// 匿名接口

	rg1 := api.Group("/user")

	{
		rg1.POST("/login", login)
		rg1.POST("/register", register)
	}

	// 需授权接口

	rg2 := api.Group("/user")

	rg2.Use(midware.Auth())

	{
		rg2.PATCH("/info", updateInfo)
		rg2.PATCH("/password", updatePassword)
	}

}
