package socket

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Socket(wsi *gin.RouterGroup) {

	rg := wsi.Group("/")

	rg.Use(midware.Auth())

	{
		rg.GET("/agent", agent)
		rg.GET("/ssh", ssh)
	}

}
