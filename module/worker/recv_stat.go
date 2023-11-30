package worker

import (
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/psutil"
)

func (pod *RecvPod) Stat(rs *SocketData) error {

	logman.Info("stat:recv", "taskId", rs.TaskId)

	err := pod.WriteJson(&SocketData{
		Method:  "Stat:resp",
		TaskId:  rs.TaskId,
		Success: true,
		Payload: psutil.Detail(true),
	})

	return err

}
