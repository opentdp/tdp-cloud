package terminal

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/terminal")

	rg.GET("/vnc", vnc)

}

func Socket(wsl *gin.RouterGroup) {

	rg := wsl.Group("/terminal")

	rg.Use(midware.Auth())

	{
		rg.GET("/ssh", ssh)
	}

}
