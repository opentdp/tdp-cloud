package worker

import (
	"log"

	"tdp-cloud/helper/psutil"
)

func (pod *SendPod) Ping() (uint, error) {

	rq := &SocketData{
		Method:  "Ping",
		TaskId:  0,
		Payload: psutil.GetSystemInfo(),
	}

	log.Println("Ping:send", "SystemInfo")

	return rq.TaskId, pod.Write(rq)

}

func (pod *RespPod) Ping(rs *SocketData) {

	log.Println("Ping:resp", rs.Payload)

}
