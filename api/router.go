package api

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"

	"tdp-cloud/api/qcloud"

	"tdp-cloud/api/secret"
	"tdp-cloud/api/ssh_key"
	"tdp-cloud/api/tat_history"
	"tdp-cloud/api/tat_script"
	"tdp-cloud/api/terminal"
	"tdp-cloud/api/user"
)

func Router(engine *gin.Engine) {

	api := engine.Group("/api")

	api.Use(midware.AbortHandle())

	{
		// qcloud api

		qcloud.Router(api)

		// direct api

		user.Router(api)
		secret.Router(api)
		ssh_key.Router(api)
		tat_history.Router(api)
		tat_script.Router(api)
	}

	// websocket

	wsl := engine.Group("/wsl")

	wsl.Use(midware.SocketPreset())

	{
		terminal.Socket(wsl)
	}

}
