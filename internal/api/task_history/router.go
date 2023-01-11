package task_history

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/internal/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/task")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/history", list)
		rg.POST("/history", create)
		rg.GET("/history/:id", detail)
		rg.PATCH("/history/:id", update)
		rg.DELETE("/history/:id", delete)
	}

}
