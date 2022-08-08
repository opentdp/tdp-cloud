package ssh_key

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/ssh/key")

	rg.Use(midware.Auth())

	{
		rg.GET("/", list)
		rg.POST("/", create)
		rg.DELETE("/:id", delete)
	}

}
