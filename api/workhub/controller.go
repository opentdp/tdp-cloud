package workhub

import (
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/opentdp/go-helper/command"
	"github.com/opentdp/go-helper/psutil"
	"github.com/opentdp/go-helper/strutil"
	"golang.org/x/net/websocket"

	"tdp-cloud/model/user"
	"tdp-cloud/module/fsadmin"
	"tdp-cloud/module/workhub"
)

type Controller struct{}

// 主机状态

func (*Controller) detail(c *gin.Context) {

	c.Set("Payload", gin.H{
		"Stat":         psutil.Detail(true),
		"MemStat":      psutil.GoMemory(),
		"NumGoroutine": runtime.NumGoroutine(),
	})

}

// 管理文件

func (*Controller) filer(c *gin.Context) {

	var rq *fsadmin.FilerParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if lst, err := fsadmin.Filer(rq); err == nil {
		c.Set("Payload", gin.H{"Items": lst})
	} else {
		c.Set("Error", err)
	}

}

// 执行脚本

func (*Controller) exec(c *gin.Context) {

	var rq *command.ExecPayload

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if ret, err := command.Exec(rq); err == nil {
		c.Set("Payload", ret)
	} else {
		c.Set("Error", err)
	}

}

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
