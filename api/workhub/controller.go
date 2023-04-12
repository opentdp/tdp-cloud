package workhub

import (
	"github.com/gin-gonic/gin"
	"github.com/open-tdp/go-helper/command"
	"github.com/spf13/cast"
	"golang.org/x/net/websocket"

	"tdp-cloud/model/user"
	"tdp-cloud/module/workhub"
)

// 节点列表

func list(c *gin.Context) {

	userId := c.GetUint("UserId")
	lst := workhub.WorkerOfUser(userId)

	c.Set("Payload", gin.H{"Items": lst})

}

// 节点状态

func detail(c *gin.Context) {

	workerId := c.Param("id")
	send := workhub.NewSender(workerId)

	if send == nil {
		c.Set("Error", "客户端已断开连接")
		return
	}

	if id, err := send.Stat(); err == nil {
		stat := workhub.WaitResponse(id, 30)
		c.Set("Payload", gin.H{"Stat": stat})
	} else {
		c.Set("Error", err)
	}

}

// 执行脚本

func exec(c *gin.Context) {

	workerId := c.Param("id")
	send := workhub.NewSender(workerId)

	if send == nil {
		c.Set("Error", "客户端已断开连接")
		return
	}

	var rq *command.ExecPayload

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if id, err := send.Exec(rq); err == nil {
		c.Set("Payload", gin.H{"Id": id})
		c.Set("Message", "下发完成")
	} else {
		c.Set("Error", err)
	}

}

// 注册节点

func register(c *gin.Context) {

	ur, err := user.Fetch(&user.FetchParam{
		AppId: c.Param("auth"),
	})

	if err != nil || ur.Id == 0 {
		c.Set("Error", "授权失败")
		return
	}

	c.Set("UserId", ur.Id)
	c.Set("MachineId", cast.ToUint(c.Param("mid")))

	// 创建 Worker 会话

	h := websocket.Handler(func(ws *websocket.Conn) {
		err := workhub.Connect(ws, &workhub.ConnectParam{
			UserId:    c.GetUint("UserId"),
			MachineId: c.GetUint("MachineId"),
		})
		c.Set("Error", err)
	})

	h.ServeHTTP(c.Writer, c.Request)

	c.Set("Payload", "连接已关闭")

}
