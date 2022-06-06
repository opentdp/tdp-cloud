package front

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed *
var vfs embed.FS

func Router(engine *gin.Engine) {

	engine.StaticFS("/ui", http.FS(vfs))

	engine.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/ui/")
	})

}
