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
		TaskId:  rq.TaskId,
		Method:  "Ping:resp",
		Payload: "OK",
	}

	return pod.Write(v)

}
