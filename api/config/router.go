package config

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/config", list)
		rg.POST("/config", create)
		rg.GET("/config/:id", detail)
		rg.PATCH("/config/:id", update)
		rg.DELETE("/config/:id", delete)

		rg.GET("/config/name/:name", detail_name)
	}

}
