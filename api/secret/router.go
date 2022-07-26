package secret

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/")

	rg.Use(midware.Auth())

	{
		rg.GET("/secret", fetchSecrets)
		rg.POST("/secret", createSecret)
		rg.DELETE("/secret/:id", deleteSecret)
	}

}
