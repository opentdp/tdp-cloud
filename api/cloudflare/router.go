package cloudflare

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

var ctrl = &Controller{}

func Router(api *gin.RouterGroup) {

	rg := api.Group("/cloudflare")

	rg.Use(midware.AuthGuard)

	{
		rg.POST("/:id", ctrl.apiProxy)
	}

}
