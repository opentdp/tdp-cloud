package workhub

import (
	"tdp-cloud/helper/command"
	"tdp-cloud/module/dborm/machine"
	"tdp-cloud/module/dborm/taskline"
)

// 更新主机

func updateMachine(worker *Worker) error {

	if worker.MachineId == 0 {
		return nil
	}

	item := &machine.UpdateParam{
		Id:         worker.MachineId,
		WorkerId:   worker.WorkerId,
		WorkerMeta: worker.WorkerMeta,
	}

	return machine.Update(item)

}

// 任务历史

func createHistory(pod *SendPod, data *command.ExecPayload) uint {

	item := &taskline.CreateParam{
		UserId:   pod.UserId,
		Subject:  "Exec: " + data.Name,
		HostName: pod.WorkerMeta.HostName,
		WorkerId: pod.WorkerId,
		Request:  data,
		Response: "",
		Status:   "Doing",
	}

	id, _ := taskline.Create(item)

	return id

}

func updateHistory(pod *RespPod, rq *SocketData) error {

	status := "Failed"
	if rq.Success {
		status = "Success"
	}

	item := &taskline.UpdateParam{
		Id:       rq.TaskId,
		UserId:   pod.UserId,
		Response: rq.Payload,
		Status:   status,
	}

	return taskline.Update(item)

}
