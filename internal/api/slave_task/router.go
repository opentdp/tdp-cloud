package slave_task

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/internal/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/slave")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/task", list)
		rg.POST("/task", create)
		rg.GET("/task/:id", detail)
		rg.PATCH("/task/:id", update)
		rg.DELETE("/task/:id", delete)
	}

}
