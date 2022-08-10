package agent

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/serve/agent"
)

func list(c *gin.Context) {

	res := agent.GetNodeList()

	c.Set("Payload", res)

}
