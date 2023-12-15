package workhub

import (
	"github.com/opentdp/go-helper/command"
	"github.com/opentdp/go-helper/socket"

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

func createHistory(worker *Worker, data *command.ExecPayload) uint {

	item := &taskline.CreateParam{
		UserId:   worker.UserId,
		Subject:  data.Name,
		HostName: worker.WorkerMeta.HostName,
		WorkerId: worker.WorkerId,
		Request:  data,
		Response: "",
		Status:   "Doing",
	}

	id, _ := taskline.Create(item)

	return id

}

func updateHistory(worker *Worker, rq *socket.PlainData) error {

	item := &taskline.UpdateParam{
		Id:     rq.TaskId,
		UserId: worker.UserId,
		Response: map[string]any{
			"Output": rq.Payload,
			"Error":  rq.Message,
		},
		Status: "Failed",
	}

	if rq.Success {
		item.Status = "Success"
	}

	return taskline.Update(item)

}
