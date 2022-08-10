package slave

import (
	"log"

	"github.com/gorilla/websocket"
)

type CommandPayload struct {
	Content          string `binding:"required"`
	Username         string `binding:"required"`
	CommandType      string `binding:"required"`
	WorkingDirectory string `binding:"required"`
	Timeout          uint   `binding:"required"`
}

func RunCommand(ws *websocket.Conn, data *CommandPayload) error {

	log.Println("RunCommand", data)

	return nil

}
