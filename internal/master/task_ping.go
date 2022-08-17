package master

import (
	"github.com/mitchellh/mapstructure"
)

func (pod *RecvPod) Ping(rq *SocketData) error {

	mapstructure.Decode(rq.Payload, &pod.SystemStat)

	v := &SocketData{
		Method:  "Ping:resp",
		TaskId:  rq.TaskId,
		Payload: "OK",
	}

	return pod.Write(v)

}
