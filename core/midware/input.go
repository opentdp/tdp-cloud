package midware

import (
	"strings"

	"tdp-cloud/core/dborm/user"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {

		input := c.Request.Header.Get("Authorization")
		field := strings.Split(input, ":")

		if len(field) != 2 {
			c.JSON(400, gin.H{"Error": "请登录后重试"})
			c.Abort()
			return
		}

		session := user.FetchSession(field[1])

		if session.UserID == 0 {
			c.JSON(400, gin.H{"Error": "会话已失效"})
			c.Abort()
			return
		}

		c.Set("KeyId", field[0])
		c.Set("UserId", session.UserID)

	}

}

func Secret() gin.HandlerFunc {

	return func(c *gin.Context) {

		keyId, _ := c.Get("KeyId")
		userId, _ := c.Get("UserId")

		secret := user.FetchSecret(keyId.(string), userId.(uint))

		if secret.ID == 0 {
			c.JSON(400, gin.H{"Error": "无法获取密钥"})
			c.Abort()
			return
		}

		c.Set("Config", [3]string{secret.SecretId, secret.SecretKey, c.Param("region")})

	}

}
