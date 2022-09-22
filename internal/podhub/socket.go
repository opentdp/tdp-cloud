package podhub

import (
	"log"
)

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

func Receiver(node *Worker) error {

	recv := &RecvPod{node}
	resp := &RespPod{node}

	for {
		var rq *SocketData

		if err := node.Read(&rq); err != nil {
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
