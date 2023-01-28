package api

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/api/config"
	"tdp-cloud/api/domain"
	"tdp-cloud/api/machine"
	"tdp-cloud/api/qcloud"
	"tdp-cloud/api/sshkey"
	"tdp-cloud/api/task_history"
	"tdp-cloud/api/task_script"
	"tdp-cloud/api/terminal"
	"tdp-cloud/api/user"
	"tdp-cloud/api/vendor"
	"tdp-cloud/api/workhub"

	"tdp-cloud/module/midware"
)

func Router(engine *gin.Engine) {

	api := engine.Group("/api")

	api.Use(midware.AbortHandle())

	{
		config.Router(api)
		domain.Router(api)
		machine.Router(api)
		qcloud.Router(api)
		vendor.Router(api)
		sshkey.Router(api)
		task_history.Router(api)
		task_script.Router(api)
		user.Router(api)
		workhub.Router(api)
	}

	// websocket interface

	wsi := engine.Group("/wsi/:appid")

	wsi.Use(midware.SocketPreset())

	{
		terminal.Socket(wsi)
		workhub.Socket(wsi)
	}

}
