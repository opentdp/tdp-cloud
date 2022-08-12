package slave

import (
	"time"

	"github.com/google/uuid"
)

func (pod *SendPod) Ping() (string, error) {

	v := &SocketData{
		TaskId:  uuid.NewString(),
		Method:  "Ping",
		Payload: time.Now().Format("2006-01-02 15:04:05"),
	}

	return v.TaskId, pod.Write(v)

}
