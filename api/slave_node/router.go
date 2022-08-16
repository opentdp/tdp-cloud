package slave_node

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/slave/node")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/", list)
		rg.POST("/exec", exec)
	}

}
