package workhub

import (
	"tdp-cloud/helper/command"
)

func (pod *SendPod) Exec(data *command.ExecPayload) (uint, error) {

	taskId := createHistory(pod, data)

	rq := &SocketData{
		Method:  "Exec",
		TaskId:  taskId,
		Payload: data,
	}

	return rq.TaskId, pod.Write(rq)

}

func (pod *RespPod) Exec(rq *SocketData) {

	updateHistory(pod, rq)

}
