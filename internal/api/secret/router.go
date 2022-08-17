package secret

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/internal/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/secret", list)
		rg.POST("/secret", create)
		rg.PATCH("/secret/:id", update)
		rg.DELETE("/secret/:id", delete)
	}

}
