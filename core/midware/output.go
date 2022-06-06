package midware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func JSON() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Next()

		// 输出错误信息

		err, _ := c.Get("Error")

		if err != nil && err != "" {
			if typeof(err) == "string" {
				err = gin.H{"message": err}
			}
			c.JSON(400, gin.H{"Error": err})
			c.Abort()
			return
		}

		// 输出请求结果

		res, _ := c.Get("Payload")

		if res != nil && res != "" {
			if typeof(res) == "string" {
				res = gin.H{"result": res}
			}
			c.JSON(200, gin.H{"Payload": res})
			c.Abort()
			return
		}

	}

}

func typeof(v interface{}) string {

	return fmt.Sprintf("%T", v)

}
