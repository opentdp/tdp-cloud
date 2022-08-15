package slave

import (
	"log"

	"tdp-cloud/core/helper"
)

func (pod *SendPod) Ping() (uint, error) {

	log.Println("Ping start")

	v := &SocketData{
		Method:  "Ping",
		TaskId:  0,
		Payload: helper.GetSystemStat(),
	}

	return v.TaskId, pod.Write(v)

}

func (pod *RespPod) Ping(rq *SocketData) {

	log.Println("Ping:resp:", rq.Payload)

}
