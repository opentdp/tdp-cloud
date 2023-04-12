package worker

import (
	"github.com/open-tdp/go-helper/logman"
	"github.com/open-tdp/go-helper/psutil"
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

	logman.Info("Ping:resp", "payload", rs.Payload)

}
