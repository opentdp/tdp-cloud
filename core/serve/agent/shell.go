package agent

import "errors"

type ShellPayload struct {
	Content          string `binding:"required"`
	Username         string `binding:"required"`
	CommandType      string `binding:"required"`
	WorkingDirectory string `binding:"required"`
	Timeout          uint   `binding:"required"`
}

func Shell(addr string, data *ShellPayload) error {

	node, ok := AgentPool[addr]

	if !ok {
		return errors.New("客户端已断开")
	}

	return node.Pod.Write(SocketData{
		Action:  "shell",
		Method:  "request",
		Payload: data,
	})

}
