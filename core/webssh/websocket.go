package webssh

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type SocketPod struct {
	*websocket.Conn
}

func NewSocketPod(w http.ResponseWriter, r *http.Request) (*SocketPod, error) {

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024 * 1024 * 10,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	ws, err := upgrader.Upgrade(w, r, nil)

	return &SocketPod{ws}, err

}

func (wsp *SocketPod) Read(p []byte) (int, error) {

	for {
		mtype, reader, err := wsp.Conn.NextReader()
		if err != nil {
			return 0, err
		}
		if mtype != websocket.TextMessage {
			continue
		}
		return reader.Read(p)
	}

}

func (wsp *SocketPod) Write(p []byte) (int, error) {

	writer, err := wsp.Conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return 0, err
	}

	defer writer.Close()
	return writer.Write(p)

}

func (wsp *SocketPod) OnClose(cb func() error) {

	wsp.Conn.SetCloseHandler(func(code int, text string) error {
		defer wsp.Conn.Close()
		return cb()
	})

}

func (wsp *SocketPod) Close() {

	wsp.Conn.Close()

}
