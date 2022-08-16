package sshkey

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/sshkey", list)
		rg.POST("/sshkey", create)
		rg.DELETE("/sshkey/:id", delete)
	}

}
