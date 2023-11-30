package workhub

import (
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/psutil"
	"github.com/opentdp/go-helper/socket"
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
	Message error
	Payload any
}

type ConnectParam struct {
	UserId    uint
	MachineId uint
}

func Connect(ws *websocket.Conn, rq *ConnectParam) error {

	pod := socket.NewWsConn(ws)
	ws.MaxPayloadBytes = 200 << 20 // 200M

	defer pod.Close()

	// 接收数据

	return Receiver(&Worker{
		pod, rq.UserId, rq.MachineId, "", "", nil, "",
	})

}

func Receiver(worker *Worker) error {

	defer DeleteWorker(worker)

	recv := &RecvPod{worker}
	resp := &RespPod{worker}

	for {
		var rq *SocketData

		if err := worker.ReadJson(&rq); err != nil {
			logman.Error("read:error", "error", err)
			return err
		}

		switch rq.Method {
		case "Register":
			recv.Register(rq)
		case "Ping":
			recv.Ping(rq)
		case "Exec:resp":
			resp.Exec(rq)
		case "Filer:resp":
			resp.Filer(rq)
		case "Stat:resp":
			resp.Stat(rq)
		default:
			logman.Warn("unknown task", "request", rq)
		}
	}

}
