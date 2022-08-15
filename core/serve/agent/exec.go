package agent

import (
	"log"

	"github.com/google/uuid"

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

func (pod *SendPod) Exec(data *ExecPayload) (string, error) {

	item := &task.CreateParam{
		UserId:   pod.UserId,
		HostId:   pod.SystemStat.HostId,
		HostName: pod.SystemStat.HostName,
		Subject:  "Exec: " + data.Name,
		Content:  helper.ToJsonString(data),
		Status:   "Doing",
		Result:   "",
	}

	task.Create(item)

	v := &SocketData{
		Method:  "Exec",
		TaskId:  uuid.NewString(),
		Payload: data,
	}

	return v.TaskId, pod.Write(v)

}

func (pod *RespPod) Exec(rq *SocketData) {

	log.Println("Ping:resp:", rq.Payload)

	item := &task.UpdateParam{
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
