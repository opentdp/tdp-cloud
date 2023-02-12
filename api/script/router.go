package script

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/script", list)
		rg.POST("/script", create)
		rg.GET("/script/:id", detail)
		rg.PATCH("/script/:id", update)
		rg.DELETE("/script/:id", delete)
	}

}
