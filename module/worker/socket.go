package worker

import (
	"log"
	"net/http"
	"time"

	"tdp-cloud/helper/psutil"
	"tdp-cloud/helper/socket"
	"tdp-cloud/helper/strutil"
	"tdp-cloud/module/workhub"
)

type SocketData = workhub.SocketData

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

	info := psutil.GetSystemInfo()
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

	go Sender(pod)
	return Receiver(pod)

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
		case "Ping:resp":
			resp.Ping(rs)
		default:
			log.Println("Task:unknown", rs)
		}
	}

}
