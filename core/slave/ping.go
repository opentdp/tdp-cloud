package slave

import (
	"log"

	"github.com/google/uuid"

	"tdp-cloud/core/helper"
)

func (pod *SendPod) Ping() (string, error) {

	log.Println("Ping start")

	v := &SocketData{
		Method:  "Ping",
		TaskId:  uuid.NewString(),
		Payload: helper.GetSystemStat(),
	}

	return v.TaskId, pod.Write(v)

}

func (pod *RespPod) Ping(rq *SocketData) {

	log.Println("Ping:resp:", rq.Payload)

}
