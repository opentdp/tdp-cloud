package workhub

import (
	"tdp-cloud/helper/command"
	"tdp-cloud/helper/logman"
)

func (pod *SendPod) Exec(data *command.ExecPayload) (uint, error) {

	logman.Info("Exec:send To", pod.WorkerMeta.HostName)

	taskId := createHistory(pod, data)

	rq := &SocketData{
		Method:  "Exec",
		TaskId:  taskId,
		Payload: data,
	}

	return rq.TaskId, pod.Write(rq)

}

func (pod *RespPod) Exec(rq *SocketData) {

	logman.Info("Exec:resp By", pod.WorkerMeta.HostName)

	updateHistory(pod, rq)

}
