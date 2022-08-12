package agent

import (
	"github.com/google/uuid"
)

type RunCommandPayload struct {
	Content          string
	Username         string
	CommandType      string
	WorkingDirectory string
	Timeout          uint
}

func (pod *SendPod) RunCommand(data *RunCommandPayload) (string, error) {

	v := &SocketData{
		TaskId:  uuid.NewString(),
		Method:  "RunCommand",
		Payload: data,
	}

	return v.TaskId, pod.Write(v)

}
