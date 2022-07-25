package member

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	// 匿名接口

	rg1 := api.Group("/member")

	{
		rg1.POST("/login", login)
		rg1.POST("/register", register)
	}

	// 需授权接口

	rg2 := api.Group("/member")

	rg2.Use(midware.Auth())

	{
		rg2.PATCH("/info", updateInfo)
		rg2.PATCH("/password", updatePassword)

		rg2.GET("/secret", fetchSecrets)
		rg2.POST("/secret", createSecret)
		rg2.DELETE("/secret/:id", deleteSecret)
	}

}
