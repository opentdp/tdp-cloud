package webssh

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type readWriter struct {
	*websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (rw *readWriter) Read(p []byte) (int, error) {

	for {
		mtype, reader, err := rw.Conn.NextReader()
		if err != nil {
			return 0, err
		}
		if mtype != websocket.TextMessage {
			continue
		}
		return reader.Read(p)
	}

}

func (rw *readWriter) Write(p []byte) (int, error) {

	writer, err := rw.Conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return 0, err
	}

	defer writer.Close()
	return writer.Write(p)

}
