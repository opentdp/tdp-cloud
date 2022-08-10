package serve

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"tdp-cloud/core/serve/agent"
	"tdp-cloud/core/socket"
)

func AgentFactory(c *gin.Context) {

	wsp, err := socket.NewJsonPod(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	defer wsp.Close()

	for {
		var rq agent.SocketData

		if wsp.Read(&rq) != nil {
			break
		}

		if rq.Action == "ping" && rq.Method == "request" {
			rs := agent.Ping(rq)
			if wsp.Write(&rs) != nil {
				break
			}
		}

	}

}

func NewSocket(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024 * 1024 * 10,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return upgrader.Upgrade(w, r, nil)

}
