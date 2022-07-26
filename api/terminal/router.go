package terminal

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/terminal")

	rg.Use(midware.Auth())
	rg.Use(midware.Secret())

}

func Socket(wsl *gin.RouterGroup) {

	rg := wsl.Group("/terminal")

	rg.Use(midware.Auth())
	rg.Use(midware.Secret())

	{
		rg.GET("/ssh", ssh)
	}

}
