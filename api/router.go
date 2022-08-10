package api

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"

	"tdp-cloud/api/qcloud"

	"tdp-cloud/api/agent"
	"tdp-cloud/api/secret"
	"tdp-cloud/api/sshkey"
	"tdp-cloud/api/tat_history"
	"tdp-cloud/api/tat_script"
	"tdp-cloud/api/user"

	"tdp-cloud/api/socket"
)

func Router(engine *gin.Engine) {

	api := engine.Group("/api")

	api.Use(midware.AbortHandle())

	{
		// qcloud api
		qcloud.Router(api)

		// direct api
		agent.Router(api)
		secret.Router(api)
		sshkey.Router(api)
		tat_history.Router(api)
		tat_script.Router(api)
		user.Router(api)
	}

	// websocket interface

	wsi := engine.Group("/wsi")

	wsi.Use(midware.SocketPreset())

	{
		socket.Socket(wsi)
	}

}
