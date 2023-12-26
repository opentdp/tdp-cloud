package workhub

import (
	"github.com/gin-gonic/gin"
	"github.com/opentdp/go-helper/strutil"
	"golang.org/x/net/websocket"

	"tdp-cloud/model/user"
	"tdp-cloud/module/workhub"
)

type Controller struct{}

// 节点列表

func (*Controller) list(c *gin.Context) {

	userId := c.GetUint("UserId")
	lst := workhub.WorkerOfUser(userId)

	c.Set("Payload", gin.H{"Items": lst})

}

// 注册节点

func (*Controller) join(c *gin.Context) {

	usr, err := user.Fetch(&user.FetchParam{
		AppId: c.Param("auth"),
	})

	if err != nil || usr.Id == 0 {
		c.Set("Error", "授权失败")
		return
	}

	rq := &workhub.ConnectParam{
		UserId:    usr.Id,
		MachineId: strutil.ToUint(c.Param("mid")),
	}

	// 创建 Worker 会话

	h := websocket.Handler(func(ws *websocket.Conn) {
		err := workhub.Connect(ws, rq)
		c.Set("Error", err)
	})

	h.ServeHTTP(c.Writer, c.Request)

	c.Set("Payload", "连接已关闭")

}
