package terminal

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/webssh"
)

func ssh(c *gin.Context) {

	option := &webssh.SSHClientOption{
		RemoteAddr: c.Query("addr"),
		User:       c.Query("user"),
		Password:   c.Query("password"),
	}

	webssh.Handle(c, option)

}
