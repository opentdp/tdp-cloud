package midware

import (
	"github.com/gin-gonic/gin"
)

func SocketHandle(c *gin.Context) {

	c.Next()

	// 输出错误信息

	if err, exists := c.Get("Error"); exists && err != nil {
		c.AbortWithError(exitCode(c, 400), newError(err))
		return
	}

}
