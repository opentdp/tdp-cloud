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

type ConnectParam struct {
	UserId    uint
	MachineId uint
}

func Connect(ws *websocket.Conn, rq *ConnectParam) error {

	conn := &socket.WsConn{Conn: ws}
	conn.MaxPayloadBytes = 200 << 20 // 200M

	worker := &Worker{
		conn, rq.UserId, rq.MachineId, "", "", nil, "",
	}

	defer worker.Close()
	defer DeleteWorker(worker)

	// 接收数据
	return Receiver(worker)

}

func Receiver(worker *Worker) error {

	recv := &RecvPod{worker}
	resp := &RespPod{worker}

	for {
		var rq *socket.PlainData
		if err := worker.ReadJson(&rq); err != nil {
			logman.Error("read:error", "error", err)
			return err
		}

		switch rq.Method {
		case "Register":
			go recv.Register(rq)
		case "Ping":
			go recv.Ping(rq)
		case "Exec:resp":
			go resp.Exec(rq)
		case "Filer:resp":
			go resp.Filer(rq)
		case "Stat:resp":
			go resp.Stat(rq)
		default:
			logman.Warn("unknown task", "request", rq)
		}
	}

}
