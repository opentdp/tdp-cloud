package user

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/")

	// 需授权接口

	rg.Use(midware.AuthGuard())
	rg.Use(midware.AdminGuard())

	{
		rg.GET("/user", list)
		rg.POST("/user", create)
		rg.GET("/user/:id", detail)
		rg.PATCH("/user/:id", update)
		rg.DELETE("/user/:id", delete)
	}

}
