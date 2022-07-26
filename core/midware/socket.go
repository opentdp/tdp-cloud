package midware

import (
	"github.com/gin-gonic/gin"
)

func SocketPreset() gin.HandlerFunc {

	return func(c *gin.Context) {

		if auth := c.Query("auth"); auth != "" {
			c.Request.Header.Set("Authorization", auth)
		}

	}

}
