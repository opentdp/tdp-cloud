package slave

import (
	"log"

	"github.com/google/uuid"
	"github.com/shirou/gopsutil/v3/host"

	"tdp-cloud/core/serve/agent"
)

type RegisterPayload agent.RegisterPayload

func (pod *SendPod) Register() (string, error) {

	log.Println("Register to server")

	data := &RegisterPayload{}

	if info, err := host.Info(); err == nil {
		data.HostId = info.HostID
		data.HostName = info.Hostname
		data.OS = info.OS
	}

	v := &SocketData{
		Method:  "Register",
		TaskId:  uuid.NewString(),
		Payload: data,
	}

	return v.TaskId, pod.Write(v)

}
