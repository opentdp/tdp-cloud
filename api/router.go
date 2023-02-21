package api

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/api/certbot"
	"tdp-cloud/api/config"
	"tdp-cloud/api/domain"
	"tdp-cloud/api/keypair"
	"tdp-cloud/api/machine"
	"tdp-cloud/api/passport"
	"tdp-cloud/api/script"
	"tdp-cloud/api/taskline"
	"tdp-cloud/api/terminal"
	"tdp-cloud/api/user"
	"tdp-cloud/api/vendor"
	"tdp-cloud/api/workhub"

	"tdp-cloud/api/alibaba"
	"tdp-cloud/api/cloudflare"
	"tdp-cloud/api/tencent"

	"tdp-cloud/module/midware"
)

func Router(engine *gin.Engine) {

	// application interface

	api := engine.Group("/api")

	api.Use(midware.OutputHandle())

	{
		certbot.Router(api)
		config.Router(api)
		domain.Router(api)
		keypair.Router(api)
		machine.Router(api)
		passport.Router(api)
		script.Router(api)
		taskline.Router(api)
		user.Router(api)
		vendor.Router(api)
		workhub.Router(api)

		alibaba.Router(api)
		cloudflare.Router(api)
		tencent.Router(api)
	}

	// websocket interface

	wsi := engine.Group("/wsi/:auth")

	wsi.Use(midware.SocketHandle())

	{
		terminal.Socket(wsi)
		workhub.Socket(wsi)
	}

}
