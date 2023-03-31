package workhub

import (
	"tdp-cloud/helper/logman"
	"tdp-cloud/helper/psutil"
	"tdp-cloud/helper/socket"
)

type SocketData struct {
	Method  string
	TaskId  uint
	Success bool
	Payload any
}

type Worker struct {
	*socket.JsonPod
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

func Daemon(worker *Worker) error {

	return Receiver(worker)

}

func Receiver(worker *Worker) error {

	recv := &RecvPod{worker}
	resp := &RespPod{worker}

	for {
		var rq *SocketData

		if err := worker.Read(&rq); err != nil {
			logman.Error("Read:error", err)
			return err
		}

		switch rq.Method {
		case "Exec:resp":
			resp.Exec(rq)
		case "Stat:resp":
			resp.Stat(rq)
		case "Ping":
			recv.Ping(rq)
		default:
			logman.Warn("Task:unknown", rq)
		}
	}

}
