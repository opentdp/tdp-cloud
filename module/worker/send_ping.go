package worker

import (
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/psutil"
	"github.com/opentdp/go-helper/socket"
)

func (pod *SendPod) Ping() (uint, error) {

	var err error

	logman.Info("ping:send")

	err = pod.WriteJson(&socket.PlainData{
		Method:  "Ping",
		TaskId:  0,
		Payload: psutil.Summary(true),
	})

	return 0, err

}

func (pod *RespPod) Ping(rs *socket.PlainData) {

	logman.Info("ping:resp", "payload", rs.Payload)

}
