package certbot

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/crontab", list)
		rg.POST("/crontab", create)
		rg.GET("/crontab/:id", detail)
		rg.PATCH("/crontab/:id", update)
		rg.DELETE("/crontab/:id", delete)
	}

}
