package sshkey

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/sshkey")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/", list)
		rg.POST("/", create)
		rg.DELETE("/:id", delete)
	}

}
