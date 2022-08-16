package tat_script

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/tat")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/script", list)
		rg.POST("/script", create)
		rg.PATCH("/script/:id", update)
		rg.DELETE("/script/:id", delete)
	}

}
