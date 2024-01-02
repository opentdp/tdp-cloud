package worker

import (
	"github.com/gin-gonic/gin"
	"github.com/opentdp/go-helper/command"

	"tdp-cloud/module/fsadmin"
	"tdp-cloud/module/workhub"
)

type Controller struct{}

// 节点状态

func (*Controller) detail(c *gin.Context) {

	workerId := c.Param("id")
	send := workhub.GetSendPod(workerId)

	if send == nil {
		c.Set("Error", "客户端连接已断开")
		return
	}

	if id, err := send.Stat(); err == nil {
		rq := workhub.WaitResponse(id, 30)
		if rq.Success {
			c.Set("Payload", rq.Payload)
		} else {
			c.Set("Error", rq.Message)
		}
	} else {
		c.Set("Error", err)
	}

}

// 管理文件

func (*Controller) filer(c *gin.Context) {

	workerId := c.Param("id")
	send := workhub.GetSendPod(workerId)

	if send == nil {
		c.Set("Error", "客户端连接已断开")
		return
	}

	var rq *fsadmin.FilerParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if id, err := send.Filer(rq); err == nil {
		rq := workhub.WaitResponse(id, 30)
		if rq.Success {
			c.Set("Payload", gin.H{"Items": rq.Payload})
		} else {
			c.Set("Error", rq.Message)
		}
	} else {
		c.Set("Error", err)
	}

}

// 执行脚本

func (*Controller) exec(c *gin.Context) {

	workerId := c.Param("id")
	send := workhub.GetSendPod(workerId)

	if send == nil {
		c.Set("Error", "客户端连接已断开")
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
