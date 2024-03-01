package midware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/opentdp/go-helper/secure"

	"tdp-cloud/cmd/args"
)

func JwtGuard(c *gin.Context) {

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
		c.Set("JwtError", "请登录后重试")
		return
	}

	// 解析并校验 Token
	claims, err := ParserToken(signToken)
	if err != nil {
		c.Set("JwtError", "会话无效，请重新登录")
		return
	}

	// 尝试解密 AppKey
	appKey, err := secure.Des3Decrypt(claims.AppKey, args.Assets.Secret)
	if err != nil {
		c.Set("JwtError", "密钥异常, 请重新注册")
		return
	}

	// 自动更新 Token
	if claims.ExpiresAt.Time.Sub(time.Now()) < 1*time.Hour {
		newToken, err := UpdateToken(signToken)
		if err == nil {
			c.Set("JwtToken", newToken)
		}
	}

	// 存储到上下文
	c.Set("AppKey", appKey)
	c.Set("UserId", claims.Id)
	c.Set("UserLevel", claims.Level)

}

func AuthGuard(c *gin.Context) {

	msg := c.GetString("JwtError")

	if msg != "" {
		c.Set("Error", gin.H{"Code": 401, "Message": msg})
		c.Set("ExitCode", 401)
		c.Abort()
	}

}

func AdminGuard(c *gin.Context) {

	id, lv := c.GetUint("UserId"), c.GetUint("UserLevel")

	if id == 0 || lv != 1 {
		c.Set("Error", gin.H{"Code": 403, "Message": "抱歉，无权进行此操作"})
		c.Set("ExitCode", 403)
		c.Abort()
	}

}
