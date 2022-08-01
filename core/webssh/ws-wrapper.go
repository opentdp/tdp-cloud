package webssh

import (
	"github.com/gorilla/websocket"
)

type wsWrapper struct {
	*websocket.Conn
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
