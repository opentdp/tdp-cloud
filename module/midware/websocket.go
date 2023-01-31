package midware

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/helper/httpd"
)

func SocketHandle() gin.HandlerFunc {

	return func(c *gin.Context) {

		if auth := c.Param("auth"); auth != "" {
			c.Request.Header.Set("Authorization", auth)
		}

		c.Next()

		// 输出错误信息

		if err, exists := c.Get("Error"); exists && err != nil {
			c.AbortWithError(errorCode(c), httpd.NewError(err))
			return
		}

	}

}
