package workhub

import (
	"github.com/open-tdp/go-helper/logman"
	"github.com/open-tdp/go-helper/psutil"
	"github.com/open-tdp/go-helper/socket"
	"golang.org/x/net/websocket"
)

type Worker struct {
	*socket.WsConn
	UserId        uint
	MachineId     uint
	CloudId       string
	WorkerId      string
	WorkerMeta    *psutil.SummaryStat
	WorkerVersion string
}

type RecvPod struct {
	*Worker
}

type RespPod struct {
	*Worker
}

type SendPod struct {
	*Worker
}

type SocketData struct {
	Method  string
	TaskId  uint
	Success bool
	Payload any
}

type ConnectParam struct {
	UserId    uint
	MachineId uint
}

func Connect(ws *websocket.Conn, rq *ConnectParam) error {

	pod := socket.NewWsConn(ws)

	defer pod.Close()

	// 接收数据

	return Receiver(&Worker{
		pod, rq.UserId, rq.MachineId, "", "", nil, "",
	})

}

func Receiver(worker *Worker) error {

	recv := &RecvPod{worker}
	resp := &RespPod{worker}

	defer delete(workerPool, worker.WorkerId)

	for {
		var rq *SocketData

		if err := worker.ReadJson(&rq); err != nil {
			logman.Error("Read:error", "error", err)
			return err
		}

		switch rq.Method {
		case "Register":
			recv.Register(rq)
		case "Ping":
			recv.Ping(rq)
		case "Exec:resp":
			resp.Exec(rq)
		case "Stat:resp":
			resp.Stat(rq)
		default:
			logman.Warn("Unknown task", "request", rq)
		}
	}

}
