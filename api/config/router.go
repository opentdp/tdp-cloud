package config

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/config")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/", list)
		rg.POST("/", create)
		rg.GET("/:name", detail)
		rg.PATCH("/:name", update)
		rg.DELETE("/:name", delete)
	}

}
