package worker

import (
	"log"
	"net/http"

	"tdp-cloud/helper/psutil"
	"tdp-cloud/helper/socket"
	"tdp-cloud/helper/strutil"
)

type SocketData struct {
	Method  string
	TaskId  uint
	Success bool
	Payload any
}

type RecvPod struct {
	*socket.JsonPod
}

type RespPod struct {
	*socket.JsonPod
}

type SendPod struct {
	*socket.JsonPod
}

func Daemon(ws string) error {

	info := psutil.Summary()
	workerId := strutil.Md5(info.HostId)

	header := http.Header{}
	header.Add("TDP-Worker-Id", workerId)
	header.Add("TDP-Worker-Meta", info.String())

	log.Println("Connecting", ws, header)
	pod, err := socket.NewJsonPodClient(ws, header)

	if err != nil {
		return err
	}

	defer pod.Close()

	go PingLoop(&SendPod{pod})

	return Receiver(pod)

}

func Receiver(pod *socket.JsonPod) error {

	recv := &RecvPod{pod}
	resp := &RespPod{pod}

	for {
		var rs *SocketData

		if err := pod.Read(&rs); err != nil {
			log.Println("Read:error", err)
			return err
		}

		switch rs.Method {
		case "Exec":
			recv.Exec(rs)
		case "Stat":
			recv.Stat(rs)
		case "Ping:resp":
			resp.Ping(rs)
		default:
			log.Println("Task:unknown", rs)
		}
	}

}
