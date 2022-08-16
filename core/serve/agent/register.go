package agent

import (
	"github.com/mitchellh/mapstructure"
)

type RegisterPayload struct {
	HostId   string
	HostName string
	OS       string
}

func (pod *RecvPod) Register(rq *SocketData, node *SlaveNode) error {

	data := &RegisterPayload{}
	mapstructure.Decode(rq.Payload, data)

	node.HostId = data.HostId
	NodePool[data.HostId] = node

	return nil

}
