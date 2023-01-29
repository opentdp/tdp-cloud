package midware

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/dborm/session"
)

func AuthGuard() gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")

		sess, err := session.Fetch(token)

		if err != nil || sess.UserId == 0 {
			c.Set("Error", "会话已失效")
			c.Abort()
			return
		}

		c.Set("UserId", sess.UserId)

	}

}
