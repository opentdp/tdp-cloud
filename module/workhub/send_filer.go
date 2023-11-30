package workhub

import (
	"fmt"
	"time"

	"github.com/opentdp/go-helper/logman"
)

type FilerPayload struct {
	Action  string
	Path    string
	Content string
}

func (pod *SendPod) Filer(data *FilerPayload) (string, error) {

	logman.Info("filer:send", "to", pod.WorkerMeta.HostName)

	taskId := uint(time.Now().UnixNano())

	err := pod.WriteJson(&SocketData{
		Method:  "Filer",
		TaskId:  taskId,
		Payload: data,
	})

	id := "filer" + fmt.Sprintf("%d", taskId)
	return id, err

}

func (pod *RespPod) Filer(rq *SocketData) {

	logman.Info("filer:resp", "from", pod.WorkerMeta.HostName)

	id := "filer" + fmt.Sprintf("%d", rq.TaskId)
	workerResp[id] = rq

}
