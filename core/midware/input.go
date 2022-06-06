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
			c.AbortWithStatusJSON(403, newError("登录后重试"))
			return
		}

		session := user.FetchSession(field[1])

		if session.UserID == 0 {
			c.AbortWithStatusJSON(403, newError("会话已失效"))
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
			c.AbortWithStatusJSON(403, newError("密钥不存在"))
			return
		}

		c.Set("Config", [3]string{secret.SecretId, secret.SecretKey, c.Param("region")})

	}

}

func newError(message string) gin.H {

	return gin.H{"Error": gin.H{"message": message}}

}
