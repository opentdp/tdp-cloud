package workhub

import (
	"time"

	"tdp-cloud/helper/logman"
)

func (pod *SendPod) Stat() (uint, error) {

	logman.Info("Stat:send To", pod.WorkerMeta.HostName)

	rq := &SocketData{
		Method: "Stat",
		TaskId: uint(time.Now().Unix()),
	}

	return rq.TaskId, pod.Write(rq)

}

func (pod *RespPod) Stat(rq *SocketData) {

	logman.Info("Stat:resp By", pod.WorkerMeta.HostName)

	workerResp[rq.TaskId] = rq.Payload

}
