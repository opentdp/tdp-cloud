package socket

import (
	"errors"

	"github.com/gin-gonic/gin"

	"tdp-cloud/core/dborm/user"
	"tdp-cloud/core/serve"
	"tdp-cloud/core/webssh"
)

func agent(c *gin.Context) {

	at := c.Param("at")

	u, err := user.Fetch(&user.FetchParam{
		AppToken: at,
	})

	if err != nil || u.Id == 0 {
		c.AbortWithError(400, errors.New("授权失败"))
		return
	}

	serve.AgentFactory(c)

}

func ssh(c *gin.Context) {

	webssh.Handle(c)

}
