package midware

import (
	"github.com/gin-gonic/gin"
)

func ExitWithJSON() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Next()

		// 输出错误信息
		if err, exists := c.Get("Error"); exists && err != nil {
			c.AbortWithStatusJSON(400, NewError(err))
			return
		}

		// 输出请求结果
		if res, exists := c.Get("Payload"); exists && res != nil {
			c.AbortWithStatusJSON(200, NewPayload(res))
			return
		}

		// 捕获异常返回
		c.AbortWithStatusJSON(500, NewError("内部错误"))

	}

}
