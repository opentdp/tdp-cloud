package midware

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"tdp-cloud/core/dborm/member"
)

func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {

		input := c.Request.Header.Get("Authorization")
		field := strings.Split(input, ":")

		if len(field) != 2 {
			c.AbortWithStatusJSON(403, NewError("登录后重试"))
			return
		}

		session := member.FetchSession(field[1])

		if session.UserId == 0 {
			c.AbortWithStatusJSON(403, NewError("会话已失效"))
			return
		}

		keyId, _ := strconv.Atoi(field[0])

		c.Set("KeyId", uint(keyId))
		c.Set("UserId", session.UserId)

	}

}

func Secret() gin.HandlerFunc {

	return func(c *gin.Context) {

		ud := GetUserdata(c)

		res, err := member.FetchSecret(ud.UserId, ud.KeyId)

		if err != nil || res.Id == 0 {
			c.AbortWithStatusJSON(403, NewError("密钥不存在"))
			return
		}

		c.Set("SecretId", res.SecretId)
		c.Set("SecretKey", res.SecretKey)

		c.Set("Region", c.Param("region"))

	}

}
