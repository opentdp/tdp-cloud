package midware

import (
	"strconv"
	"strings"

	"tdp-cloud/core/dborm/user"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {

		input := c.Request.Header.Get("Authorization")
		field := strings.Split(input, ":")

		if len(field) != 2 {
			c.AbortWithStatusJSON(403, NewError("登录后重试"))
			return
		}

		session := user.FetchSession(field[1])

		if session.UserId == 0 {
			c.AbortWithStatusJSON(403, NewError("会话已失效"))
			return
		}

		keyId, _ := strconv.Atoi(field[0])

		c.Set("KeyId", keyId)
		c.Set("UserId", session.UserId)

	}

}

func Secret() gin.HandlerFunc {

	return func(c *gin.Context) {

		ud := GetUserdata(c)

		secret := user.FetchSecret(ud.UserId, ud.KeyId)

		if secret.Id == 0 {
			c.AbortWithStatusJSON(403, NewError("密钥不存在"))
			return
		}

		c.Set("Region", c.Param("region"))
		c.Set("SecretId", secret.SecretId)
		c.Set("SecretKey", secret.SecretKey)

	}

}
