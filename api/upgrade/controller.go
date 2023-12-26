package upgrade

import (
	"github.com/gin-gonic/gin"
	"github.com/opentdp/go-helper/upgrade"

	"tdp-cloud/cmd/args"
)

type Controller struct{}

// 检查升级

func (*Controller) check(c *gin.Context) {

	rq := &upgrade.RequesParam{
		Server:  args.UpdateUrl,
		Version: args.Version,
	}

	if res, err := upgrade.CheckVersion(rq); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 执行升级

func (*Controller) apply(c *gin.Context) {

	rq := &upgrade.RequesParam{
		Server:  args.UpdateUrl,
		Version: args.Version,
	}

	if err := upgrade.Apply(rq); err == nil {
		if err := upgrade.Restart(); err != nil {
			c.Set("Error", err)
		}
		c.Set("Message", "更新完成")
	} else {
		c.Set("Error", err)
	}

}
