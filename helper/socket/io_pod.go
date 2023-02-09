package socket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type IOPod struct {
	*websocket.Conn
}

func NewIOPod(w http.ResponseWriter, r *http.Request) (*IOPod, error) {

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024 * 1024 * 10,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)

	return &IOPod{conn}, err

}

func (pod *IOPod) Read(p []byte) (int, error) {

	for {
		mtype, reader, err := pod.Conn.NextReader()
		if err != nil {
			return 0, err
		}
		if mtype != websocket.TextMessage {
			continue
		}
		return reader.Read(p)
	}

}

func (pod *IOPod) Write(p []byte) (int, error) {

	writer, err := pod.Conn.NextWriter(websocket.TextMessage)

	if err != nil {
		return 0, err
	}

	defer writer.Close()
	return writer.Write(p)

}

func (pod *IOPod) OnClose(cb func() error) {

	pod.Conn.SetCloseHandler(func(code int, text string) error {
		defer pod.Conn.Close()
		return cb()
	})

}

func (pod *IOPod) Close() {

	pod.Conn.Close()

}
