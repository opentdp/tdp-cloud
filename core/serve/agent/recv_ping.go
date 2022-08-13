package agent

import (
	"github.com/mitchellh/mapstructure"
)

func (pod *RecvPod) Ping(rq *SocketData) error {

	addr := pod.Conn.RemoteAddr().String()

	if node, ok := AgentPool[addr]; ok {
		mapstructure.Decode(rq.Payload, &node.Stat)
	}

	v := &SocketData{
		TaskId:  rq.TaskId,
		Method:  "Ping:end",
		Payload: "OK",
	}

	return pod.Write(v)

}
