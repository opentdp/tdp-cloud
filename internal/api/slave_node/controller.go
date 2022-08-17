package slave_node

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/internal/slaver"
)

func list(c *gin.Context) {

	userId := c.GetUint("UserId")

	res := slaver.NodesOfUser(userId)

	c.Set("Payload", res)

}

type execParam struct {
	HostId  string
	Payload slaver.ExecPayload
}

func exec(c *gin.Context) {

	var rq *execParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	send := slaver.NewSender(rq.HostId)

	if send == nil {
		c.Set("Error", "客户端已断开连接")
		return
	}

	if id, err := send.Exec(&rq.Payload); err == nil {
		c.Set("Payload", map[string]any{
			"Message": "命令下发完成",
			"TaskId":  id,
		})
	} else {
		c.Set("Error", err)
	}

}
