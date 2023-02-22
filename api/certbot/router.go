package certbot

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/certbot", list)
		rg.POST("/certbot", create)
		rg.GET("/certbot/:id", detail)
		rg.PATCH("/certbot/:id", update)
		rg.DELETE("/certbot/:id", delete)
	}

}
