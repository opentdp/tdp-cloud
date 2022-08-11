package agent

import "errors"

func Pong(addr string, data *interface{}) error {

	node, ok := AgentPool[addr]

	if !ok {
		return errors.New("客户端已断开")
	}

	v := &SocketData{
		Action:  "pong",
		Method:  "response",
		Payload: data,
	}

	return node.Pod.Write(v)

}
