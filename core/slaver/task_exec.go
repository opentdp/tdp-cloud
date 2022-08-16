package slaver

import (
	"tdp-cloud/core/helper"

	task "tdp-cloud/core/dborm/slave_task"
)

type ExecPayload struct {
	Name             string
	CommandType      string
	Content          string
	Username         string
	WorkingDirectory string
	Timeout          uint
}

func (pod *SendPod) Exec(data *ExecPayload) (uint, error) {

	item := &task.CreateParam{
		UserId:   pod.UserId,
		HostId:   pod.SystemStat.HostId,
		HostName: pod.SystemStat.HostName,
		Subject:  "Exec: " + data.Name,
		Content:  helper.ToJsonString(data),
		Status:   "Doing",
		Result:   "",
	}

	taskId, _ := task.Create(item)

	v := &SocketData{
		Method:  "Exec",
		TaskId:  taskId,
		Payload: data,
	}

	return v.TaskId, pod.Write(v)

}

func (pod *RespPod) Exec(rq *SocketData) {

	item := &task.UpdateParam{
		Id:     rq.TaskId,
		UserId: pod.UserId,
		Result: helper.ToJsonString(rq.Payload),
	}

	if rq.Success {
		item.Status = "Success"
	} else {
		item.Status = "Failed"
	}

	task.Update(item)

}
