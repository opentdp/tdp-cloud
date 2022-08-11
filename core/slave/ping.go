package slave

import (
	"time"

	"tdp-cloud/core/socket"
)

func Ping(pod *socket.JsonPod) error {

	v := SocketData{
		Action:  "ping",
		Method:  "request",
		Payload: time.Now().Format("2006-01-02 15:04:05"),
	}

	return pod.Write(v)

}
