package machine

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

var ctrl = &Controller{}

func Router(api *gin.RouterGroup) {

	rg := api.Group("/machine")

	rg.Use(midware.AuthGuard)

	{
		rg.POST("/list", ctrl.list)
		rg.POST("/create", ctrl.create)
		rg.POST("/detail", ctrl.detail)
		rg.POST("/update", ctrl.update)
		rg.POST("/delete", ctrl.delete)
	}

}
