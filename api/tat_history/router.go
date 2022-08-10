package tat_history

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/tat/history")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/", list)
		rg.POST("/", create)
		rg.PATCH("/:id", update)
		rg.DELETE("/:id", delete)
	}

}
