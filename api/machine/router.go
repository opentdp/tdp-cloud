package machine

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/machine", list)
		rg.POST("/machine", create)
		rg.GET("/machine/:id", detail)
		rg.PATCH("/machine/:id", update)
		rg.DELETE("/machine/:id", delete)
	}

}
