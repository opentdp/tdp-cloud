package slaver

import (
	"log"
)

type RecvPod struct {
	*SlaveNode
}
type RespPod struct {
	*SlaveNode
}
type SendPod struct {
	*SlaveNode
}

type SocketData struct {
	Method  string
	TaskId  uint
	Success bool
	Payload any
}

func Receiver(node *SlaveNode) error {

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
