package task_script

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/task")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/script", list)
		rg.POST("/script", create)
		rg.PATCH("/script/:id", update)
		rg.DELETE("/script/:id", delete)
	}

}
