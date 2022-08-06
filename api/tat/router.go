package tat

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/tat")

	rg.Use(midware.Auth())

	{
		rg.POST("/history/", create_history)
		rg.GET("/history/", list_history)
		rg.PATCH("/history/:id", update_history)
		rg.DELETE("/history/:id", delete_history)
		rg.GET("/", list)
		rg.POST("/", create)
		rg.PATCH("/:id", update)
		rg.DELETE("/:id", delete)
	}

}
