package workhub

import (
	"log"

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
	UserId     uint
	MachineId  uint
	WorkerId   string
	WorkerMeta *psutil.SummaryStat
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
			log.Println("Read:error", err)
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
			log.Println("Task:unknown", rq)
		}
	}

}
