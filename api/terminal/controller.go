package terminal

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"tdp-cloud/helper/webssh"
	"tdp-cloud/module/model/keypair"
)

func ssh(c *gin.Context) {

	// 获取 SSH 参数

	var rq *webssh.SSHClientOption

	if err := c.ShouldBindQuery(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if id := cast.ToUint(c.Param("id")); id > 0 {
		kp, err := keypair.Fetch(&keypair.FetchParam{
			Id:       id,
			UserId:   c.GetUint("UserId"),
			StoreKey: c.GetString("AppKey"),
		})
		if err != nil || kp.Id == 0 {
			c.Set("Error", "密钥不存在")
			return
		}
		rq.PrivateKey = kp.PrivateKey
	}

	// 创建 SSH 连接

	err := webssh.Connect(&webssh.ConnectParam{
		Request: c.Request,
		Writer:  c.Writer,
		Option:  rq,
	})

	if err != nil {
		c.Set("Error", err)
		return
	}

}
