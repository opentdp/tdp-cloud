package worker

import (
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/psutil"
)

func (pod *SendPod) Ping() (uint, error) {

	logman.Info("ping:send")

	stat := psutil.Summary(true)

	err := pod.WriteJson(&SocketData{
		Method:  "Ping",
		TaskId:  0,
		Payload: stat,
	})

	return 0, err

}

func (pod *RespPod) Ping(rs *SocketData) {

	logman.Info("ping:resp", "payload", rs.Payload)

}
