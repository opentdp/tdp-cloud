package terminal

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

var ctrl = &Controller{}

func Socket(wsi *gin.RouterGroup) {

	rg := wsi.Group("/terminal")

	rg.Use(midware.AuthGuard)

	{
		rg.GET("/ssh/:id", ctrl.ssh)
	}

}
