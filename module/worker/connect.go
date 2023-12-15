package worker

import (
	"os"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/socket"

	"tdp-cloud/cmd/args"
)

type RecvPod struct {
	*socket.WsConn
}

type RespPod struct {
	*socket.WsConn
}

type SendPod struct {
	*socket.WsConn
}

func Connect() error {

	url := args.Worker.Remote
	hostname, _ := os.Hostname() // 获取主机名
	conn, err := socket.NewWsClient(url, "", "tdp://"+hostname)
	if err != nil {
		return err
	}

	conn.MaxPayloadBytes = 200 << 20 // 200M
	defer conn.Close()

	// 注册节点
	send := &SendPod{conn}
	go send.Register()

	// 接收数据
	return Receiver(conn)

}

func Receiver(conn *socket.WsConn) error {

	recv := &RecvPod{conn}
	resp := &RespPod{conn}

	for {
		var rq *socket.PlainData
		if err := conn.ReadJson(&rq); err != nil {
			logman.Error("read json failed", "error", err)
			return err
		}

		switch rq.Method {
		case "Exec":
			go recv.Exec(rq)
		case "Stat":
			go recv.Stat(rq)
		case "Filer":
			go recv.Filer(rq)
		case "Ping:resp":
			go resp.Ping(rq)
		case "Register:resp":
			go resp.Register(rq)
		default:
			logman.Warn("unknown task", "request", rq)
		}
	}

}
