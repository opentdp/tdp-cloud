package agent

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/serve/agent"
)

func list(c *gin.Context) {

	res := agent.GetNodeList()

	c.Set("Payload", res)

}

type commandParam struct {
	Addr    string
	Payload agent.RunCommandPayload
}

func runCommand(c *gin.Context) {

	var rq *commandParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	taskId, err := agent.SendRunCommand(rq.Addr, &rq.Payload)

	//TODO: 返回的TaskId需要入库，以便异步接收回调结果

	if err == nil {
		c.Set("Payload", "命令下发完成，TaskId："+taskId)
	} else {
		c.Set("Error", err)
	}

}
