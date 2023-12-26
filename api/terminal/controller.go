package terminal

import (
	"github.com/gin-gonic/gin"
	"github.com/opentdp/go-helper/strutil"
	"github.com/opentdp/go-helper/webssh"
	"golang.org/x/net/websocket"

	"tdp-cloud/model/keypair"
)

type Controller struct{}

func (*Controller) ssh(c *gin.Context) {

	// 获取 SSH 参数

	var rq *webssh.SSHClientOption

	if err := c.ShouldBindQuery(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if id := strutil.ToUint(c.Param("id")); id > 0 {
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

	h := websocket.Handler(func(ws *websocket.Conn) {
		err := webssh.Connect(ws, rq)
		c.Set("Error", err)
	})

	h.ServeHTTP(c.Writer, c.Request)

	c.Set("Payload", "连接已关闭")

}
