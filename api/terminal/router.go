package terminal

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Socket(wsi *gin.RouterGroup) {

	rg := wsi.Group("/terminal")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/ssh", ssh)
	}

}
