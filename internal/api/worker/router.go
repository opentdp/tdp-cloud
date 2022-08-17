package worker

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/internal/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/worker", list)
		rg.POST("/worker/exec", exec)
	}

}
