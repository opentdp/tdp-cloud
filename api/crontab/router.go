package crontab

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	ctrl := &Controller{}

	rg := api.Group("/crontab")

	rg.Use(midware.AuthGuard)

	{
		rg.POST("/list", ctrl.list)
		rg.POST("/create", ctrl.create)
		rg.POST("/detail", ctrl.detail)
		rg.POST("/update", ctrl.update)
		rg.POST("/delete", ctrl.delete)
	}

}
