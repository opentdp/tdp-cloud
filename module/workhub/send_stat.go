package workhub

import (
	"time"
)

func (pod *SendPod) Stat() (uint, error) {

	rq := &SocketData{
		Method: "Stat",
		TaskId: uint(time.Now().Unix()),
	}

	return rq.TaskId, pod.Write(rq)

}

func (pod *RespPod) Stat(rq *SocketData) {

	workerResp[rq.TaskId] = rq.Payload

}
