package midware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"tdp-cloud/cmd/args"
	"tdp-cloud/helper/strutil"
)

func AuthGuard() gin.HandlerFunc {

	return func(c *gin.Context) {

		signToken := ""

		// 取回已签名 Token
		authcode := c.GetHeader("Authorization")
		parts := strings.SplitN(authcode, " ", 2)
		if len(parts) == 2 && parts[0] == "Bearer" {
			signToken = parts[1]
		} else {
			signToken = c.Param("auth")
		}

		// 找不到有效 Token
		if signToken == "" {
			c.Set("Error", gin.H{"Code": 401, "Message": "请登录后重试"})
			c.Set("ExitCode", 401)
			c.Abort()
			return
		}

		// 解析并校验 Token
		claims, err := ParserToken(signToken)
		if err != nil {
			c.Set("Error", gin.H{"Code": 401, "Message": "会话无效，请重新登录"})
			c.Set("ExitCode", 401)
			c.Abort()
			return
		}

		// 尝试解密 AppKey
		appKey, err := strutil.Des3Decrypt(claims.AppKey, args.Dataset.Secret)
		if err != nil {
			c.Set("Error", gin.H{"Code": 401, "Message": "密钥异常, 请重新注册"})
			c.Set("ExitCode", 401)
			c.Abort()
			return
		}

		// 存储到上下文
		c.Set("AppKey", appKey)
		c.Set("UserId", claims.Id)
		c.Set("UserLevel", claims.Level)

	}

}

func AdminGuard() gin.HandlerFunc {

	return func(c *gin.Context) {

		id, lv := c.GetUint("UserId"), c.GetUint("UserLevel")

		if id == 0 || lv != 1 {
			c.Set("Error", gin.H{"Code": 403, "Message": "抱歉，无权进行此操作"})
			c.Set("ExitCode", 403)
			c.Abort()
			return
		}

	}

}
