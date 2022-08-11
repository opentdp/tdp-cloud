package slave

import (
	"time"

	"github.com/google/uuid"

	"tdp-cloud/core/socket"
)

func SendPing(pod *socket.JsonPod) (string, error) {

	v := &SocketData{
		TaskId:  uuid.New().String(),
		Method:  "Ping",
		Payload: time.Now().Format("2006-01-02 15:04:05"),
	}

	return v.TaskId, pod.Write(v)

}
