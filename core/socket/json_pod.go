package socket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type JsonPod struct {
	*websocket.Conn
}

func NewJsonPod(w http.ResponseWriter, r *http.Request) (*JsonPod, error) {

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024 * 1024 * 10,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	ws, err := upgrader.Upgrade(w, r, nil)

	return &JsonPod{ws}, err

}

func (pod *JsonPod) Read(v *interface{}) error {

	return pod.Conn.ReadJSON(v)

}

func (pod *JsonPod) Write(v *interface{}) error {

	return pod.Conn.WriteJSON(v)

}

func (pod *JsonPod) OnClose(cb func() error) {

	pod.Conn.SetCloseHandler(func(code int, text string) error {
		defer pod.Conn.Close()
		return cb()
	})

}

func (pod *JsonPod) Close() {

	pod.Conn.Close()

}
