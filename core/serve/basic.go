package serve

import (
	"os"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func Create(addr string, mids ...func(*gin.Engine)) {

	if os.Getenv("IS_DEBUG") == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine = gin.Default()

	for _, mid := range mids {
		mid(engine)
	}

	Listen(addr)

}
