package terminal

import (
	"log"

	"github.com/gin-gonic/gin"

	"tdp-cloud/core/webssh"
)

func ssh(c *gin.Context) {

	log.Println("Webssh - Connecting")

	webssh.Handle(c)

	log.Println("Webssh - Disconnected")

}
