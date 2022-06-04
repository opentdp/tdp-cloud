package user

import (
	"github.com/gin-gonic/gin"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/user")

	{
		rg.POST("/login", Login)
		rg.POST("/register", Register)
	}

}
