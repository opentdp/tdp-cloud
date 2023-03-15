package workhub

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"tdp-cloud/helper/command"
	"tdp-cloud/helper/psutil"
	"tdp-cloud/module/model/user"
	"tdp-cloud/module/workhub"
)

// 主机信息

func host(c *gin.Context) {

	info := psutil.Detail()

	c.Set("Payload", gin.H{"Stat": info})

}

// 主机IP

func hostIp(c *gin.Context) {

	ip := psutil.PublicIpAddress(false)

	c.Set("Payload", gin.H{"Ip": ip})

}

// 节点列表

func list(c *gin.Context) {

	userId := c.GetUint("UserId")
	lst := workhub.WorkerOfUser(userId)

	c.Set("Payload", gin.H{"Items": lst})

}

// 获取状态

func stat(c *gin.Context) {

	workerId := c.Param("id")
	send := workhub.NewSender(workerId)

	if send == nil {
		c.Set("Error", "客户端已断开连接")
		return
	}

	if id, err := send.Stat(); err == nil {
		res := workhub.WaitResponse(id, 30)
		c.Set("Payload", res)
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

	u, err := user.Fetch(&user.FetchParam{
		AppId: c.Param("auth"),
	})

	if err != nil || u.Id == 0 {
		c.Set("Error", "授权失败")
		return
	}

	c.Set("UserId", u.Id)
	c.Set("MachineId", cast.ToUint(c.Param("mid")))

	if err := workhub.Register(c); err != nil {
		c.Set("Error", err)
		return
	}

}
