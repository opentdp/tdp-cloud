package agent

import "errors"

func RecvPing(addr string, rq *SocketData) error {

	node, ok := AgentPool[addr]

	if !ok {
		return errors.New("客户端已断开")
	}

	v := &SocketData{
		TaskId:  rq.TaskId,
		Method:  "Ping:end",
		Payload: rq.Payload,
	}

	return node.Pod.Write(v)

}
