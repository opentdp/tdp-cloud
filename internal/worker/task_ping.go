package worker

import (
	"log"

	"tdp-cloud/internal/helper"
)

func (pod *SendPod) Ping() (uint, error) {

	v := &SocketData{
		Method:  "Ping",
		TaskId:  0,
		Payload: helper.GetSystemStat(),
	}

	log.Println("Ping:send", "SystemStat")

	return v.TaskId, pod.Write(v)

}

func (pod *RespPod) Ping(rq *SocketData) {

	log.Println("Ping:resp", rq.Payload)

}
