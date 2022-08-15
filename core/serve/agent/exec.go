package agent

import (
	"log"

	"github.com/google/uuid"

	"tdp-cloud/core/helper"

	history "tdp-cloud/core/dborm/tat_history"
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

	v := &SocketData{
		Method:  "Exec",
		TaskId:  uuid.NewString(),
		Payload: data,
	}

	createHistory(pod.UserId, v.TaskId, data)

	return v.TaskId, pod.Write(v)

}

func (pod *RespPod) Exec(rq *SocketData) {

	log.Println("Ping:resp:", rq.Payload)

	item := &history.UpdateParam{
		InvocationId:         rq.TaskId,
		InvocationResultJson: helper.ToJsonString(rq.Payload),
	}

	if rq.Success {
		item.InvocationStatus = "Success"
	} else {
		item.InvocationStatus = "Failed"
	}

	history.Update(item)

}

/////

func createHistory(userId uint, taskId string, data *ExecPayload) error {

	return history.Create(&history.CreateParam{
		UserId:       userId,
		KeyId:        uint(0),
		Name:         data.Name,
		InvocationId: taskId,
		Region:       "agent",
	})

}
