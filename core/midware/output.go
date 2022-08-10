package midware

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/helper"
)

func AbortHandle() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Next()

		// 输出错误信息
		if err, exists := c.Get("Error"); exists && err != nil {
			c.AbortWithStatusJSON(400, helper.NewMessage(err))
			return
		}

		// 输出请求结果
		if res, exists := c.Get("Payload"); exists && res != nil {
			c.AbortWithStatusJSON(200, helper.NewPayload(res))
			return
		}

		// 输出HTML内容
		if res, exists := c.Get("HTML"); exists && res != nil {
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.String(200, res.(string))
			c.Abort()
			return
		}

		// 捕获异常返回
		c.AbortWithStatusJSON(500, helper.NewMessage("内部错误"))

	}

}
