package worktask

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/internal/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/worktask", list)
		rg.POST("/worktask", create)
		rg.GET("/worktask/:id", detail)
		rg.PATCH("/worktask/:id", update)
		rg.DELETE("/worktask/:id", delete)
	}

}
