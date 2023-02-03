package taskline

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/taskline", list)
		rg.POST("/taskline", create)
		rg.GET("/taskline/:id", detail)
		rg.PATCH("/taskline/:id", update)
		rg.DELETE("/taskline/:id", delete)
	}

}
