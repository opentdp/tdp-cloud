package worker

import (
	"tdp-cloud/cmd/args"
	"tdp-cloud/helper/logman"
	"tdp-cloud/helper/socket"
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

type SocketData struct {
	Method  string
	TaskId  uint
	Success bool
	Payload any
}

func Connect() error {

	url := args.Worker.Remote
	pod, err := socket.NewWsClient(url, "", "http://localhost")

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
		var rq *SocketData

		if err := pod.ReadJson(&rq); err != nil {
			logman.Error("Read json failed", "Error", err)
			return err
		}

		switch rq.Method {
		case "Exec":
			recv.Exec(rq)
		case "Stat":
			recv.Stat(rq)
		case "Ping:resp":
			resp.Ping(rq)
		case "Register:resp":
			resp.Register(rq)
		default:
			logman.Warn("Unknown task", "SocketData", rq)
		}
	}

}
