package agent

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/serve/agent"
)

func list(c *gin.Context) {

	userId := c.GetUint("UserId")

	res := agent.NodesOfUser(userId)

	c.Set("Payload", res)

}

type execParam struct {
	Addr    string
	Payload agent.ExecPayload
}

func exec(c *gin.Context) {

	var rq *execParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	pod := agent.NewSendPod(rq.Addr)

	if pod == nil {
		c.Set("Error", "客户端已断开连接")
		return
	}

	taskId, err := pod.Exec(&rq.Payload)

	if err == nil {
		c.Set("Payload", "命令下发完成，TaskId："+taskId)
	} else {
		c.Set("Error", err)
	}

}
