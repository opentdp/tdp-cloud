package midware

import (
	"github.com/gin-gonic/gin"
)

func SocketHandle() gin.HandlerFunc {

	return func(c *gin.Context) {

		if auth := c.Param("auth"); auth != "" {
			c.Request.Header.Set("Authorization", auth)
		}

	}

}
