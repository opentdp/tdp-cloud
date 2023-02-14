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

func Daemon(woker *Worker) error {

	return Receiver(woker)

}

func Receiver(woker *Worker) error {

	recv := &RecvPod{woker}
	resp := &RespPod{woker}

	for {
		var rq *SocketData

		if err := woker.Read(&rq); err != nil {
			log.Println("Read:error", err)
			return err
		}

		switch rq.Method {
		case "Exec:resp":
			resp.Exec(rq)
		case "Ping":
			recv.Ping(rq)
		default:
			log.Println("Task:unknown", rq)
		}
	}

}
