package agent

import (
	"github.com/mitchellh/mapstructure"
)

func (pod *RecvPod) Ping(rq *SocketData) error {

	addr := pod.Conn.RemoteAddr().String()

	if node, ok := NodePool[addr]; ok {
		mapstructure.Decode(rq.Payload, &node.SystemStat)
	}

	v := &SocketData{
		Method:  "Ping:resp",
		TaskId:  rq.TaskId,
		Payload: "OK",
	}

	return pod.Write(v)

}
