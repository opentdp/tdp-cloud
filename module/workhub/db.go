package workhub

import (
	"github.com/open-tdp/go-helper/command"

	"tdp-cloud/model/machine"
	"tdp-cloud/model/taskline"
)

// 更新主机

func updateMachine(worker *Worker) error {

	// 尝试查找 MachineId
	if worker.MachineId == 0 && worker.CloudId != "" {
		item, err := machine.Fetch(&machine.FetchParam{
			CloudId: worker.CloudId,
		})
		if err == nil && item.Id > 0 {
			worker.MachineId = item.Id
		}
	}

	// 忽略更新没有注册的主机
	if worker.MachineId == 0 {
		return nil
	}

	item := &machine.UpdateParam{
		Id:         worker.MachineId,
		CloudId:    worker.CloudId,
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
