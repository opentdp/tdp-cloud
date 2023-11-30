package workhub

import (
	"fmt"
	"time"

	"github.com/opentdp/go-helper/logman"
)

func (pod *SendPod) Stat() (string, error) {

	logman.Info("stat:send", "to", pod.WorkerMeta.HostName)

	taskId := uint(time.Now().UnixNano())

	err := pod.WriteJson(&SocketData{
		Method: "Stat",
		TaskId: taskId,
	})

	id := "stat" + fmt.Sprintf("%d", taskId)
	return id, err

}

func (pod *RespPod) Stat(rq *SocketData) {

	logman.Info("stat:resp", "from", pod.WorkerMeta.HostName)

	id := "stat" + fmt.Sprintf("%d", rq.TaskId)
	workerResp[id] = rq

}
