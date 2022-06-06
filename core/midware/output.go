package midware

import (
	"github.com/gin-gonic/gin"
)

func JSON() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Next()

		err, _ := c.Get("Error")

		if err != nil && err != "" {
			c.JSON(400, gin.H{"Error": err})
			c.Abort()
			return
		}

		res, _ := c.Get("Payload")

		if res != nil && res != "" {
			c.JSON(200, gin.H{"Payload": res})
			c.Abort()
			return
		}

	}

}
