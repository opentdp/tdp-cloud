package webssh

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type wsWrapper struct {
	*websocket.Conn
}

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (wsw *wsWrapper) Read(p []byte) (int, error) {

	for {
		mtype, reader, err := wsw.Conn.NextReader()
		if err != nil {
			return 0, err
		}
		if mtype != websocket.TextMessage {
			continue
		}
		return reader.Read(p)
	}

}

func (wsw *wsWrapper) Write(p []byte) (int, error) {

	writer, err := wsw.Conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return 0, err
	}

	defer writer.Close()
	return writer.Write(p)

}

func (wsw *wsWrapper) SetCloseHandler(cb func() error) {

	wsw.Conn.SetCloseHandler(func(code int, text string) error {
		wsw.Conn.Close()
		return cb()
	})

}
