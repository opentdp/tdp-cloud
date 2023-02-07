package keypair

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/keypair", list)
		rg.POST("/keypair", create)
		rg.DELETE("/keypair/:id", delete)
	}

}
