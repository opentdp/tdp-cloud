package agent

import (
	"errors"

	"github.com/google/uuid"
)

type RunCommandPayload struct {
	Content          string
	Username         string
	CommandType      string
	WorkingDirectory string
	Timeout          uint
}

func SendRunCommand(addr string, data *RunCommandPayload) (string, error) {

	node, ok := AgentPool[addr]

	if !ok {
		return "", errors.New("客户端已断开")
	}

	v := &SocketData{
		TaskId:  uuid.New().String(),
		Method:  "RunCommand",
		Payload: data,
	}

	return v.TaskId, node.Pod.Write(v)

}
