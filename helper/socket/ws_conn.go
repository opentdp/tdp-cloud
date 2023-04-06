package socket

import (
	"golang.org/x/net/websocket"
)

type WsConn struct {
	*websocket.Conn
}

func (pod *WsConn) Read(v []byte) error {
	return websocket.Message.Receive(pod.Conn, v)
}

func (pod *WsConn) ReadJson(v any) error {
	return websocket.JSON.Receive(pod.Conn, v)
}

func (pod *WsConn) Write(p []byte) error {
	return websocket.Message.Send(pod.Conn, p)
}

func (pod *WsConn) WriteJson(v any) error {
	return websocket.JSON.Send(pod.Conn, v)
}

func (pod *WsConn) Close() error {
	return pod.Conn.Close()
}

func (pod *WsConn) Die(r string) {
	pod.Write([]byte(r))
	pod.Conn.Close()
}

// 创建连接

func NewWsConn(ws *websocket.Conn) *WsConn {

	return &WsConn{ws}

}

func NewWsClient(u, p, o string) (*WsConn, error) {

	ws, err := websocket.Dial(u, p, o)
	if err != nil {
		return nil, err
	}

	return &WsConn{ws}, nil

}