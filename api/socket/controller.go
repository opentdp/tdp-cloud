package socket

import (
	"log"

	"github.com/gin-gonic/gin"

	"tdp-cloud/core/serve"
	"tdp-cloud/core/webssh"
)

func agent(c *gin.Context) {

	log.Println("agent - Connecting")

	serve.AgentFactory(c)

	log.Println("agent - Disconnected")

}

func ssh(c *gin.Context) {

	log.Println("webssh - Connecting")

	webssh.Handle(c)

	log.Println("webssh - Disconnected")

}
