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
		rg.GET("/config/:name", detail)
		rg.PATCH("/config/:name", update)
		rg.DELETE("/config/:name", delete)
	}

}
