package workhub

import (
	"github.com/mitchellh/mapstructure"
	"github.com/opentdp/go-helper/logman"
)

func (pod *RecvPod) Register(rq *SocketData) error {

	logman.Info("register:recv", "from", pod.Conn.RemoteAddr())

	// 注册主机

	worker := pod.Worker

	if err := mapstructure.Decode(rq.Payload, worker); err != nil {
		pod.Die("register:error " + err.Error())
	}

	if worker.WorkerId == "" {
		pod.Die("register:error WorkerId is empty")
	}

	if err := updateMachine(worker); err != nil {
		pod.Die("register:error " + err.Error())
	}

	workerPool[worker.WorkerId] = worker

	// 返回结果

	err := pod.WriteJson(&SocketData{
		Method:  "Register:resp",
		TaskId:  rq.TaskId,
		Success: true,
		Payload: "OK, Id: " + worker.WorkerId,
	})

	return err

}
