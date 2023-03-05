package workhub

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/workhub")

	rg.Use(midware.AuthGuard())

	{
		rg.POST("/host", host)
		rg.POST("/list", list)
		rg.POST("/stat/:id", stat)
		rg.POST("/exec/:id", exec)
	}

}

func Socket(wsi *gin.RouterGroup) {

	rg := wsi.Group("/")

	{
		rg.GET("/workhub", register)
		rg.GET("/workhub/:mid", register)
	}

}
