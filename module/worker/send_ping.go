package worker

import (
	"log"

	"tdp-cloud/helper/psutil"
)

func (pod *SendPod) Ping() (uint, error) {

	rq := &SocketData{
		Method:  "Ping",
		TaskId:  0,
		Payload: psutil.Summary(),
	}

	log.Println("Ping:send", "SummaryStat")

	return rq.TaskId, pod.Write(rq)

}

func (pod *RespPod) Ping(rs *SocketData) {

	log.Println("Ping:resp", rs.Payload)

}
