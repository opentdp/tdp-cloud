package worker

import (
	"net/http"

	"tdp-cloud/cmd/args"
	"tdp-cloud/helper/crypto"
	"tdp-cloud/helper/logman"
	"tdp-cloud/helper/psutil"
	"tdp-cloud/helper/socket"
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

func Daemon() error {

	ws := args.Worker.Remote

	info := psutil.Summary()
	cloudId := psutil.CloudInstanceId()
	workerId := crypto.Md5ToString(info.HostId)

	header := http.Header{}
	header.Add("TDP-Cloud-Id", cloudId)
	header.Add("TDP-Worker-Id", workerId)
	header.Add("TDP-Worker-Meta", info.String())

	logman.Warn("Connecting", ws, header)
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
			logman.Error("Read:error", err)
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
			logman.Warn("Task:unknown", rs)
		}
	}

}
