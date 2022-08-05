package terminal

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"tdp-cloud/core/webssh"
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

	log.Println("Webssh - Connecting")

	webssh.Handle(c)

	log.Println("Webssh - Disconnected")

}
