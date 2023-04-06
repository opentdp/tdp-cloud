package socket

import (
	"io"
	"net"

	"golang.org/x/net/websocket"
)

type TcpRelayParam struct {
	targetAddr string
	binaryMode bool
}

func TcpRelay(ws *websocket.Conn, rq *TcpRelayParam) error {

	// 接受客户端连接

	defer ws.Close()

	if rq.binaryMode {
		ws.PayloadType = websocket.BinaryFrame
	}

	// 连接远程服务器

	conn, err := net.Dial("tcp", rq.targetAddr)

	if err != nil {
		return err
	}

	defer conn.Close()

	// 保持连接并转发数据

	ch := make(chan error)
	defer close(ch)

	go ioCopy(conn, ws, ch)
	go ioCopy(ws, conn, ch)

	return <-ch

}

func ioCopy(dst io.Writer, src io.Reader, ch chan<- error) {

	_, err := io.Copy(dst, src)
	ch <- err

}
