package socket

import (
	"errors"

	"github.com/gin-gonic/gin"

	"tdp-cloud/helper/webssh"
	"tdp-cloud/internal/dborm/user"
	"tdp-cloud/internal/workhub"
)

func worker(c *gin.Context) {

	u, err := user.Fetch(&user.FetchParam{
		AppId: c.Param("appid"),
	})

	if err != nil || u.Id == 0 {
		c.AbortWithError(400, errors.New("授权失败"))
		return
	}

	c.Set("UserId", u.Id)

	workhub.Register(c)

}

func ssh(c *gin.Context) {

	// 获取 SSH 参数

	var option *webssh.SSHClientOption

	if err := c.ShouldBindQuery(&option); err != nil {
		c.AbortWithError(500, err)
		return
	}

	// 创建 SSH 连接

	err := webssh.Connect(&webssh.ConnectParam{
		Request: c.Request,
		Writer:  c.Writer,
		Option:  option,
	})

	if err != nil {
		c.AbortWithError(500, err)
		return
	}

}
