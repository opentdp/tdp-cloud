package api

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"

	"tdp-cloud/api/cloud"

	"tdp-cloud/api/secret"
	"tdp-cloud/api/tat"
	"tdp-cloud/api/terminal"
	"tdp-cloud/api/user"
)

func Router(engine *gin.Engine) {

	api := engine.Group("/api")

	{
		api.Use(midware.ExitWithJSON())

		// cloud api

		cloud.Router(api)

		// local api

		user.Router(api)
		secret.Router(api)
		tat.Router(api)
	}

	// websocket

	wsl := engine.Group("/wsl")

	{
		wsl.Use(midware.SocketPreset())

		terminal.Socket(wsl)
	}

}
