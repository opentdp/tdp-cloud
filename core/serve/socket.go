package serve

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func UseSocekt(c *gin.Context) {

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	_, message, err := ws.ReadMessage()
	if nil != err {
		fmt.Println(err.Error())
	}

	ws.WriteMessage(websocket.TextMessage, message)

}
