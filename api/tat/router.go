package tat

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/tat")

	rg.Use(midware.Auth())

	rg.GET("list", listTAT)
	rg.GET("info/:id", infoTAT)
	rg.POST("create", createTAT)
	rg.POST("update", updateTAT)
	rg.DELETE("delete/:id", deleteTAT)
}
