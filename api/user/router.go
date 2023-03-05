package user

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/user")

	// 管理员接口

	rg.Use(midware.AuthGuard())
	rg.Use(midware.AdminGuard())

	{
		rg.POST("/list", list)
		rg.POST("/create", create)
		rg.POST("/detail", detail)
		rg.POST("/update", update)
		rg.POST("/delete", delete)
	}

}
