package midware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSON() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Next()

		if err, _ := c.Get("Error"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err,
			})
			return
		}

		if res, _ := c.Get("Payload"); res != nil {
			c.JSON(http.StatusOK, gin.H{
				"Payload": res,
			})
			return
		}

	}

}
