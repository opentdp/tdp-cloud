package midware

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/model/session"
)

func AuthGuard() gin.HandlerFunc {

	return func(c *gin.Context) {

		sess, err := session.Fetch(&session.FetchParam{
			Token: c.GetHeader("Authorization"),
		})

		if err != nil || sess.UserId == 0 {
			c.Set("Error", "会话已失效")
			c.Abort()
			return
		}

		c.Set("UserId", sess.UserId)
		c.Set("UserLevel", sess.UserLevel)

	}

}

func AdminGuard() gin.HandlerFunc {

	return func(c *gin.Context) {

		id, lv := c.GetUint("UserId"), c.GetUint("UserLevel")

		if id == 0 || lv != 1 {
			c.Set("Error", "无权限")
			c.Abort()
			return
		}

	}

}
