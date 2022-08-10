package serve

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func AgentFactory(c *gin.Context) {

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	_, message, err := ws.ReadMessage()
	if nil != err {
		log.Println(err.Error())
	}

	ws.WriteMessage(websocket.TextMessage, message)

}
