package workhub

import (
	"time"

	"github.com/opentdp/go-helper/logman"
)

func (pod *SendPod) Stat() (uint, error) {

	logman.Info("stat:send", "to", pod.WorkerMeta.HostName)

	taskId := uint(time.Now().Unix())

	err := pod.WriteJson(&SocketData{
		Method: "Stat",
		TaskId: taskId,
	})

	return taskId, err

}

func (pod *RespPod) Stat(rq *SocketData) {

	logman.Info("stat:resp", "from", pod.WorkerMeta.HostName)

	workerResp[rq.TaskId] = rq.Payload

}
