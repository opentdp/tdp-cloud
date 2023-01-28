package worker

import (
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/host"

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

func Daemon(ws string) {

	info, _ := host.Info()
	workerId := strutil.Md5(info.HostID)

	args := []string{
		"WorkerId=" + workerId,
		"HostName=" + url.QueryEscape(info.Hostname),
		"OSType=" + info.OS,
	}

	ws += "?" + strings.Join(args, "&")

	log.Println("Connecting", ws)
	pod, err := socket.NewJsonPodClient(ws)

	if err != nil {
		return
	}

	defer pod.Close()

	go Sender(pod)
	Receiver(pod)

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
