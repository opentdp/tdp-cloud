package midware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSON() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Next()

		err, _ := c.Get("Error")

		if err != nil && err != "" {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err})
			return
		}

		res, _ := c.Get("Payload")

		if res != nil && res != "" {
			c.JSON(http.StatusOK, gin.H{"Payload": res})
			return
		}

	}

}
