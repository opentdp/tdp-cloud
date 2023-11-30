package workhub

import (
	"time"

	"github.com/opentdp/go-helper/logman"
)

type FilerPayload struct {
	Action  string
	Path    string
	Content string
}

func (pod *SendPod) Filer(data *FilerPayload) (uint, error) {

	logman.Info("filer:send", "to", pod.WorkerMeta.HostName)

	taskId := uint(time.Now().Unix())

	err := pod.WriteJson(&SocketData{
		Method:  "Filer",
		TaskId:  taskId,
		Payload: data,
	})

	return taskId, err

}

func (pod *RespPod) Filer(rq *SocketData) {

	logman.Info("filer:resp", "from", pod.WorkerMeta.HostName)

	workerResp[rq.TaskId] = rq.Payload

}
