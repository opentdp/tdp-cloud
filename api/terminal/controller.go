package terminal

import (
	"io/ioutil"
	"net/http"
	"tdp-cloud/core/webssh"

	"github.com/gin-gonic/gin"
)

func vnc(c *gin.Context) {

	resp, err := http.Get("https://img.qcloud.com/qcloud/app/active_vnc/index.html")
	if err != nil {
		c.Set("Error", "获取资源失败")
		return
	}

	if res, err := ioutil.ReadAll(resp.Body); err == nil {
		c.Set("HTML", string(res))
	} else {
		c.Set("Error", err)
	}

}

func ssh(c *gin.Context) {

	var rq webssh.SSHClientOption

	if err := c.ShouldBindQuery(&rq); err != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	webssh.Handle(c, &rq)

}
