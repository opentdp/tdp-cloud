package agent

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/serve/agent"
)

func list(c *gin.Context) {

	res := agent.GetNodeList()

	c.Set("Payload", res)

}

type shellParam struct {
	Addr    string
	Payload agent.ShellPayload
}

func shell(c *gin.Context) {

	var rq shellParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	if err := agent.Shell(rq.Addr, &rq.Payload); err == nil {
		c.Set("Payload", "命令下发完成")
	} else {
		c.Set("Error", err)
	}

}
