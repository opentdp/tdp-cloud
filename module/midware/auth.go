package midware

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/dborm/session"
	"tdp-cloud/module/dborm/user"
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

	}

}

func AdminGuard() gin.HandlerFunc {

	return func(c *gin.Context) {

		rq := &user.FetchParam{
			Id: c.GetUint("UserId"),
		}

		user, err := user.Fetch(rq)

		if err != nil || user.Level != 1 {
			c.Set("Error", "无权限")
			c.Abort()
			return
		}

	}

}
