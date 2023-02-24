package workhub

import (
	"github.com/mitchellh/mapstructure"

	"tdp-cloud/helper/logman"
)

func (pod *RecvPod) Ping(rq *SocketData) error {

	logman.Info("Ping:recv By", pod.WorkerMeta.HostName)

	mapstructure.Decode(rq.Payload, &pod.WorkerMeta)

	rs := &SocketData{
		Method:  "Ping:resp",
		TaskId:  rq.TaskId,
		Payload: "OK",
	}

	return pod.Write(rs)

}
