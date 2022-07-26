package tat

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/")

	rg.Use(midware.Auth())

	{
		rg.GET("/tat", listTAT)
		rg.GET("/tat/:id", infoTAT)
		rg.POST("/tat", createTAT)
		rg.PATCH("/tat/:id", updateTAT)
		rg.DELETE("/tat/:id", deleteTAT)
	}

}
