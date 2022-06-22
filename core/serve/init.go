package serve

import (
	"os"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func Create(addr string, useRoute func(*gin.Engine)) {

	if os.Getenv("IS_DEBUG") == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine = gin.Default()

	useRoute(engine)

	Listen(addr)

}
