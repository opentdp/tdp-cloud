package midware

import (
	"tdp-cloud/core/dborm"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {

		var session dborm.Session

		token := c.Request.Header.Get("Token")
		dborm.Db.First(&session, "token = ?", token)

		c.Set("UserID", session.UserID)

		c.Next()

	}

}
