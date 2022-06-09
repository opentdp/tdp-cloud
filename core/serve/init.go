package serve

import (
	"os"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func Create() *gin.Engine {

	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine = gin.Default()

	return engine

}
