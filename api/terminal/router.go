package terminal

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Socket(wsl *gin.RouterGroup) {

	rg := wsl.Group("/terminal")

	rg.Use(midware.Auth())

	{
		rg.GET("/ssh", ssh)
	}

}
