package agent

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/agent")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/node", list)
		rg.POST("/command", runCommand)
	}

}
