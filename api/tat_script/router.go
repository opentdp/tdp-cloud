package tat_script

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/tat/script")

	rg.Use(midware.Auth())

	{
		rg.GET("/", list)
		rg.POST("/", create)
		rg.PATCH("/:id", update)
		rg.DELETE("/:id", delete)
	}

}
