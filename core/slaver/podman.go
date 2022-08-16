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

func Receiver(node *SlaveNode) {

	recv := &RecvPod{node}
	resp := &RespPod{node}

	for {
		var rq *SocketData

		if node.Read(&rq) != nil {
			break
		}

		switch rq.Method {
		case "Exec:resp":
			resp.Exec(rq)
		case "Register":
			recv.Register(rq, node)
		case "Ping":
			recv.Ping(rq)
		default:
			log.Println("Unknown task:", rq)
		}
	}

}

func NewSendPod(hostId string) *SendPod {

	if node, ok := NodePool[hostId]; ok {
		return &SendPod{node}
	}

	return nil

}
