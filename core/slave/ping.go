package slave

import (
	"time"

	"github.com/gorilla/websocket"
)

func Ping(ws *websocket.Conn) error {

	v := SocketData{
		Action:  "ping",
		Method:  "request",
		Payload: time.Now().Format("2006-01-02 15:04:05"),
	}

	return ws.WriteJSON(v)

}
