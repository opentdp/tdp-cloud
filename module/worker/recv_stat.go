package worker

import (
	"github.com/open-tdp/go-helper/logman"
	"github.com/open-tdp/go-helper/psutil"
)

func (pod *RecvPod) Stat(rs *SocketData) error {

	logman.Info("stat:recv", "taskId", rs.TaskId)

	err := pod.WriteJson(&SocketData{
		Method:  "Stat:resp",
		TaskId:  rs.TaskId,
		Payload: psutil.Detail(true),
	})

	return err

}
