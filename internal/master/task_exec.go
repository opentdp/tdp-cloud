package master

import (
	"tdp-cloud/helper/json"

	"tdp-cloud/internal/dborm/worktask"
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

	item := &worktask.CreateParam{
		UserId:   pod.UserId,
		HostId:   pod.SystemStat.HostId,
		HostName: pod.SystemStat.HostName,
		Subject:  "Exec: " + data.Name,
		Status:   "Doing",
		Request:  json.ToString(data),
		Response: "",
	}

	taskId, _ := worktask.Create(item)

	// 发送给节点执行该任务

	v := &SocketData{
		Method:  "Exec",
		TaskId:  taskId,
		Payload: data,
	}

	return v.TaskId, pod.Write(v)

}

func (pod *RespPod) Exec(rq *SocketData) {

	item := &worktask.UpdateParam{
		Id:       rq.TaskId,
		UserId:   pod.UserId,
		Response: json.ToString(rq.Payload),
	}

	if rq.Success {
		item.Status = "Success"
	} else {
		item.Status = "Failed"
	}

	worktask.Update(item)

}
