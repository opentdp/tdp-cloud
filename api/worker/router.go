package worker

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	ctrl := &Controller{}

	rg := api.Group("/worker/:id")

	rg.Use(midware.AuthGuard)

	{
		rg.POST("/detail", ctrl.detail)
		rg.POST("/exec", ctrl.exec)
		rg.POST("/filer", ctrl.filer)
	}

}
