package slave

import (
	"log"
	"time"

	"tdp-cloud/core/slaver"
	"tdp-cloud/core/socket"
)

type RecvPod struct {
	*socket.JsonPod
}

type RespPod struct {
	*socket.JsonPod
}

type SendPod struct {
	*socket.JsonPod
}

type SocketData slaver.SocketData

func Receiver(pod *socket.JsonPod) {

	recv := &RecvPod{pod}
	resp := &RespPod{pod}

	for {
		var rq *SocketData

		if pod.Read(&rq) != nil {
			break
		}

		switch rq.Method {
		case "Exec":
			recv.Exec(rq)
		case "Ping:resp":
			resp.Ping(rq)
		default:
			log.Println("Unknown task:", rq)
		}
	}

}

func Sender(pod *socket.JsonPod) {

	send := &SendPod{pod}

	if _, err := send.Register(); err != nil {
		return
	}

	for {

		if _, err := send.Ping(); err != nil {
			log.Println(err)
			break
		}

		time.Sleep(time.Second * 15)

	}

}
