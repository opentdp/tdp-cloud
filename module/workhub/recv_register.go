package workhub

import (
	"github.com/mitchellh/mapstructure"
	"github.com/open-tdp/go-helper/logman"
)

func (pod *RecvPod) Register(rq *SocketData) error {

	logman.Info("Register:recv", "from", pod.Conn.RemoteAddr())

	// 注册主机

	worker := pod.Worker
	workerPool[worker.WorkerId] = worker

	if err := mapstructure.Decode(rq.Payload, worker); err != nil {
		pod.Die("Register:error " + err.Error())
	}

	if err := updateMachine(worker); err != nil {
		pod.Die("Register:error " + err.Error())
	}

	// 返回结果

	err := pod.WriteJson(&SocketData{
		Method:  "Register:resp",
		TaskId:  rq.TaskId,
		Payload: "OK",
	})

	return err

}
