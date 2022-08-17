package worker

import (
	"log"
	"time"

	"github.com/shirou/gopsutil/v3/host"

	"tdp-cloud/helper/socket"
	"tdp-cloud/internal/master"
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

type SocketData master.SocketData

func Register(url string) {

	if hostId, err := host.HostID(); err == nil {
		url += "?HostId=" + hostId
	}

	pod, err := socket.NewJsonPodClient(url)

	if err != nil {
		return
	}

	defer pod.Close()

	go Sender(pod)
	Receiver(pod)

}

func Receiver(pod *socket.JsonPod) error {

	recv := &RecvPod{pod}
	resp := &RespPod{pod}

	for {
		var rq *SocketData

		if err := pod.Read(&rq); err != nil {
			log.Println("Read:error", err)
			return err
		}

		switch rq.Method {
		case "Exec":
			recv.Exec(rq)
		case "Ping:resp":
			resp.Ping(rq)
		default:
			log.Println("Task:unknown", rq)
		}
	}

}

func Sender(pod *socket.JsonPod) error {

	send := &SendPod{pod}

	for {

		if _, err := send.Ping(); err != nil {
			log.Println("Send:error", err)
			return err
		}

		time.Sleep(time.Second * 15)

	}

}
