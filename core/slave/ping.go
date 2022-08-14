package slave

import (
	"log"

	"github.com/google/uuid"

	"tdp-cloud/core/helper"
)

func (pod *SendPod) Ping() (string, error) {

	v := &SocketData{
		TaskId:  uuid.NewString(),
		Method:  "Ping",
		Payload: helper.GetSystemStat(),
	}

	return v.TaskId, pod.Write(v)

}

func (pod *RespPod) Ping(rq *SocketData) {

	log.Println("Ping:resp:", rq.Payload)

}
