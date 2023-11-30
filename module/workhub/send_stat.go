package workhub

import (
	"fmt"
	"time"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/socket"
)

func (pod *SendPod) Stat() (string, error) {

	var (
		err    error
		taskId = uint(time.Now().UnixNano())
	)

	logman.Info("stat:send", "to", pod.WorkerMeta.HostName)

	err = pod.WriteJson(&socket.PlainData{
		Method: "Stat",
		TaskId: taskId,
	})

	id := "stat" + fmt.Sprintf("%d", taskId)
	return id, err

}

func (pod *RespPod) Stat(rs *socket.PlainData) {

	logman.Info("stat:resp", "from", pod.WorkerMeta.HostName)

	id := "stat" + fmt.Sprintf("%d", rs.TaskId)
	workerResp[id] = rs

}
