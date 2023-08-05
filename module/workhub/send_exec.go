package workhub

import (
	"github.com/opentdp/go-helper/command"
	"github.com/opentdp/go-helper/logman"
)

func (pod *SendPod) Exec(data *command.ExecPayload) (uint, error) {

	logman.Info("exec:send", "to", pod.WorkerMeta.HostName)

	taskId := createHistory(pod, data)

	err := pod.WriteJson(&SocketData{
		Method:  "Exec",
		TaskId:  taskId,
		Payload: data,
	})

	return taskId, err

}

func (pod *RespPod) Exec(rq *SocketData) {

	logman.Info("exec:resp", "from", pod.WorkerMeta.HostName)

	updateHistory(pod, rq)

}
