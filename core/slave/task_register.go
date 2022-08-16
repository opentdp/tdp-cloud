package slave

import (
	"log"

	"github.com/shirou/gopsutil/v3/host"

	"tdp-cloud/core/slaver"
)

type RegisterPayload slaver.RegisterPayload

func (pod *SendPod) Register() (uint, error) {

	log.Println("Register to server...")

	data := &RegisterPayload{}

	if info, err := host.Info(); err == nil {
		data.HostId = info.HostID
		data.HostName = info.Hostname
		data.OS = info.OS
	}

	v := &SocketData{
		Method:  "Register",
		TaskId:  0,
		Payload: data,
	}

	return v.TaskId, pod.Write(v)

}
