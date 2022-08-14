package midware

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"tdp-cloud/core/dborm/session"
)

func AuthGuard() gin.HandlerFunc {

	return func(c *gin.Context) {

		input := c.Request.Header.Get("Authorization")
		field := strings.Split(input, ":")

		if len(field) != 2 {
			c.Set("Error", "登录后重试")
			c.Abort()
			return
		}

		sess, err := session.Fetch(field[1])

		if err != nil || sess.UserId == 0 {
			c.Set("Error", "会话已失效")
			c.Abort()
			return
		}

		keyId, _ := strconv.Atoi(field[0])

		c.Set("KeyId", uint(keyId))
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
