package agent

import (
	"github.com/mitchellh/mapstructure"
)

type RegisterPayload struct {
	HostId   string
	HostName string
	OS       string
}

func (pod *RecvPod) Register(rq *SocketData) error {

	data := &RegisterPayload{}
	mapstructure.Decode(rq.Payload, data)

	return nil

}
