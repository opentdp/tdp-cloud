package workhub

import (
	"log"
	"tdp-cloud/helper/command"
)

func (pod *SendPod) Exec(data *command.ExecPayload) (uint, error) {

	log.Println("Exec:send To", pod.WorkerMeta.HostName)

	taskId := createHistory(pod, data)

	rq := &SocketData{
		Method:  "Exec",
		TaskId:  taskId,
		Payload: data,
	}

	return rq.TaskId, pod.Write(rq)

}

func (pod *RespPod) Exec(rq *SocketData) {

	log.Println("Exec:resp By", pod.WorkerMeta.HostName)

	updateHistory(pod, rq)

}
