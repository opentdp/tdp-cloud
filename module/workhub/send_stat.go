package workhub

import (
	"time"

	"tdp-cloud/helper/logman"
)

func (pod *SendPod) Stat() (uint, error) {

	logman.Info("Stat:send", "to", pod.WorkerMeta.HostName)

	taskId := uint(time.Now().Unix())

	err := pod.WriteJson(&SocketData{
		Method: "Stat",
		TaskId: taskId,
	})

	return taskId, err

}

func (pod *RespPod) Stat(rq *SocketData) {

	logman.Info("Stat:resp", "from", pod.WorkerMeta.HostName)

	workerResp[rq.TaskId] = rq.Payload

}
