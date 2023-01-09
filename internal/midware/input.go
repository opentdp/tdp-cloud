package midware

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/internal/dborm/session"
)

func AuthGuard() gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.Request.Header.Get("Authorization")

		sess, err := session.Fetch(token)

		if err != nil || sess.UserId == 0 {
			c.Set("Error", "会话已失效")
			c.Abort()
			return
		}

		c.Set("UserId", sess.UserId)

	}

}

func SocketPreset() gin.HandlerFunc {

	return func(c *gin.Context) {

		if auth := c.Param("auth"); auth != "" {
			c.Request.Header.Set("Authorization", auth)
		}

	}

}
