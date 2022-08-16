package slave_task

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/slave/task")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/", list)
		rg.POST("/", create)
		rg.GET("/:id", detail)
		rg.PATCH("/:id", update)
		rg.DELETE("/:id", delete)
	}

}
