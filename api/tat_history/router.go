package tat_history

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/tat")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/history", list)
		rg.POST("/history", create)
		rg.PATCH("/history/:id", update)
		rg.DELETE("/history/:id", delete)
	}

}
