package slave

import (
	"tdp-cloud/core/helper"

	"github.com/google/uuid"
)

func (pod *SendPod) Ping() (string, error) {

	v := &SocketData{
		TaskId:  uuid.NewString(),
		Method:  "Ping",
		Payload: helper.GetSystemStat(),
	}

	return v.TaskId, pod.Write(v)

}
