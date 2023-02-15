package workhub

import (
	"log"

	"github.com/mitchellh/mapstructure"
)

func (pod *RecvPod) Ping(rq *SocketData) error {

	log.Println("Ping:recv By", pod.WorkerMeta.HostName)

	mapstructure.Decode(rq.Payload, &pod.WorkerMeta)

	rs := &SocketData{
		Method:  "Ping:resp",
		TaskId:  rq.TaskId,
		Payload: "OK",
	}

	return pod.Write(rs)

}
