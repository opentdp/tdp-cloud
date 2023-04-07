package workhub

import (
	"tdp-cloud/helper/command"
	"tdp-cloud/helper/logman"
)

func (pod *SendPod) Exec(data *command.ExecPayload) (uint, error) {

	logman.Info("Exec:send", "To", pod.WorkerMeta.HostName)

	taskId := createHistory(pod, data)

	err := pod.WriteJson(&SocketData{
		Method:  "Exec",
		TaskId:  taskId,
		Payload: data,
	})

	return taskId, err

}

func (pod *RespPod) Exec(rq *SocketData) {

	logman.Info("Exec:resp", "From", pod.WorkerMeta.HostName)

	updateHistory(pod, rq)

}
