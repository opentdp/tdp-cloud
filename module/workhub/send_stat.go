package workhub

import (
	"log"
	"time"
)

func (pod *SendPod) Stat() (uint, error) {

	log.Println("Stat:send To", pod.WorkerMeta.HostName)

	rq := &SocketData{
		Method: "Stat",
		TaskId: uint(time.Now().Unix()),
	}

	return rq.TaskId, pod.Write(rq)

}

func (pod *RespPod) Stat(rq *SocketData) {

	log.Println("Stat:resp By", pod.WorkerMeta.HostName)

	workerResp[rq.TaskId] = rq.Payload

}
