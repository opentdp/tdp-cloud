package api

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/internal/midware"

	"tdp-cloud/internal/api/qcloud"

	"tdp-cloud/internal/api/config"
	"tdp-cloud/internal/api/secret"
	"tdp-cloud/internal/api/slave_node"
	"tdp-cloud/internal/api/slave_task"
	"tdp-cloud/internal/api/sshkey"
	"tdp-cloud/internal/api/tat_history"
	"tdp-cloud/internal/api/tat_script"
	"tdp-cloud/internal/api/user"

	"tdp-cloud/internal/api/socket"
)

func Router(engine *gin.Engine) {

	api := engine.Group("/api")

	api.Use(midware.AbortHandle())

	{
		// qcloud api
		qcloud.Router(api)

		// direct api
		config.Router(api)
		secret.Router(api)
		slave_node.Router(api)
		slave_task.Router(api)
		sshkey.Router(api)
		tat_history.Router(api)
		tat_script.Router(api)
		user.Router(api)
	}

	// websocket interface

	wsi := engine.Group("/wsi/:auth")

	wsi.Use(midware.SocketPreset())

	{
		socket.Socket(wsi)
	}

}
