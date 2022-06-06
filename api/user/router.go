package user

import (
	"tdp-cloud/core/midware"

	"github.com/gin-gonic/gin"
)

func Router(api *gin.RouterGroup) {

	rg1 := api.Group("/user")

	{
		rg1.POST("/login", login)
		rg1.POST("/register", register)
	}

	rg2 := api.Group("/user")

	rg2.Use(midware.Auth())

	{
		rg2.GET("/secret", fetchSecrets)
		rg2.POST("/secret", createSecret)
		rg2.DELETE("/secret/:id", deleteSecret)
	}

}
