package terminal

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/webssh"
)

func ssh(c *gin.Context) {

	confing := &webssh.WebSSHConfig{
		RemoteAddr: c.Query("addr"),
		User:       c.Query("user"),
		Password:   c.Query("password"),
		AuthModel:  webssh.PASSWORD,
	}

	wsh := webssh.NewWebSSH(confing)

	wsh.ServeConn(c)

}
