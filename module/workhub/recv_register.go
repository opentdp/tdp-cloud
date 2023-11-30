package workhub

import (
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/socket"
)

func (pod *RecvPod) Register(rq *socket.PlainData) error {

	var (
		err    error
		worker = pod.Worker
	)

	logman.Info("register:recv", "from", pod.Conn.RemoteAddr())

	// 注册主机

	if err = rq.GetPayload(worker); err != nil {
		pod.Die("register:error " + err.Error())
	}

	if worker.WorkerId == "" {
		pod.Die("register:error WorkerId is empty")
	}

	if err = updateMachine(worker); err != nil {
		pod.Die("register:error " + err.Error())
	}

	workerPool[worker.WorkerId] = worker

	// 返回结果

	err = pod.WriteJson(&socket.PlainData{
		Method:  "Register:resp",
		TaskId:  rq.TaskId,
		Success: true,
		Payload: "OK, Id: " + worker.WorkerId,
	})

	return err

}
