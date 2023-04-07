package worker

import (
	"tdp-cloud/helper/logman"
	"tdp-cloud/helper/psutil"
)

func (pod *SendPod) Ping() (uint, error) {

	logman.Info("Ping:send")

	stat := psutil.Summary(true)

	err := pod.WriteJson(&SocketData{
		Method:  "Ping",
		TaskId:  0,
		Payload: stat,
	})

	return 0, err

}

func (pod *RespPod) Ping(rs *SocketData) {

	logman.Info("Ping:resp", "Payload", rs.Payload)

}
