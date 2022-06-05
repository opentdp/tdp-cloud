package user

import (
	"github.com/gin-gonic/gin"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/user")

	{
		rg.POST("/login", login)
		rg.POST("/register", register)

		rg.GET("/secret", fetchSecrets)
		rg.POST("/secret", createSecret)
		rg.DELETE("/secret/:id", deleteSecret)
	}

}
