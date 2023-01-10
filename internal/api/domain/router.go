package domain

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/internal/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/domain", list)
		rg.POST("/domain", create)
		rg.GET("/domain/:id", detail)
		rg.PATCH("/domain/:id", update)
		rg.DELETE("/domain/:id", delete)
	}

}
