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
	pod, err := socket.NewWsClient(url, "", "tdp://"+hostname)

	if err != nil {
		return err
	}

	defer pod.Close()

	// 注册节点

	send := &SendPod{pod}
	go send.Register()

	// 接收数据

	return Receiver(pod)

}

func Receiver(pod *socket.WsConn) error {

	recv := &RecvPod{pod}
	resp := &RespPod{pod}

	for {
		var rq *socket.PlainData

		if err := pod.ReadJson(&rq); err != nil {
			logman.Error("read json failed", "error", err)
			return err
		}

		switch rq.Method {
		case "Exec":
			recv.Exec(rq)
		case "Stat":
			recv.Stat(rq)
		case "Filer":
			recv.Filer(rq)
		case "Ping:resp":
			resp.Ping(rq)
		case "Register:resp":
			go resp.Register(rq)
		default:
			logman.Warn("unknown task", "request", rq)
		}
	}

}
