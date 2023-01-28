package workhub

import (
	"github.com/spf13/cast"

	"tdp-cloud/module/dborm/machine"
	history "tdp-cloud/module/dborm/task_history"
)

// 绑定主机

func bindMachine(id, workerId string) error {

	item := &machine.UpdateParam{
		Id:       cast.ToUint(id),
		WorkerId: workerId,
	}

	return machine.Update(item)

}

// 任务历史

func createHistory(pod *SendPod, data *ExecPayload) uint {

	item := &history.CreateParam{
		UserId:   pod.UserId,
		Subject:  "Exec: " + data.Name,
		HostName: pod.SystemStat.HostName,
		WorkerId: pod.WorkerId,
		Request:  data,
		Response: "",
		Status:   "Doing",
	}

	id, _ := history.Create(item)

	return id

}

func updateHistory(pod *RespPod, rq *SocketData) error {

	status := "Failed"
	if rq.Success {
		status = "Success"
	}

	item := &history.UpdateParam{
		Id:       rq.TaskId,
		UserId:   pod.UserId,
		Response: rq.Payload,
		Status:   status,
	}

	return history.Update(item)

}
