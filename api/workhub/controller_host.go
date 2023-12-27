package workhub

import (
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/opentdp/go-helper/command"
	"github.com/opentdp/go-helper/psutil"

	"tdp-cloud/module/fsadmin"
)

type HostController struct{}

// 主机状态

func (*HostController) detail(c *gin.Context) {

	c.Set("Payload", gin.H{
		"Stat":         psutil.Detail(true),
		"MemStat":      psutil.GoMemory(),
		"NumGoroutine": runtime.NumGoroutine(),
	})

}

// 管理文件

func (*HostController) filer(c *gin.Context) {

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

func (*HostController) exec(c *gin.Context) {

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
