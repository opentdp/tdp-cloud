package slave

import (
	"log"

	"tdp-cloud/core/socket"
)

type CommandPayload struct {
	Content          string `binding:"required"`
	Username         string `binding:"required"`
	CommandType      string `binding:"required"`
	WorkingDirectory string `binding:"required"`
	Timeout          uint   `binding:"required"`
}

func RunCommand(pod *socket.JsonPod, data *CommandPayload) error {

	log.Println("RunCommand", data)

	return nil

}
