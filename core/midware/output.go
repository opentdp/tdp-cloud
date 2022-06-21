package midware

import (
	"github.com/gin-gonic/gin"
)

func JSON() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Next()

		// 输出错误信息

		if errAny, exists := c.Get("Error"); exists {
			if err, ok := errAny.(string); ok {
				c.AbortWithStatusJSON(400, NewError(err))
				return
			}

			if err, ok := errAny.(error); ok {
				c.AbortWithStatusJSON(400, NewError(err.Error()))
				return
			}
		}

		// 输出请求结果

		if res, exists := c.Get("Payload"); exists {
			c.AbortWithStatusJSON(200, gin.H{"Payload": res})
			return
		}
	}

}
