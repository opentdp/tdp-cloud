package worker

import (
	"tdp-cloud/helper/logman"
	"tdp-cloud/helper/psutil"
)

func (pod *RecvPod) Stat(rs *SocketData) error {

	logman.Info("Stat:recv", "TaskId", rs.TaskId)

	err := pod.WriteJson(&SocketData{
		Method:  "Stat:resp",
		TaskId:  rs.TaskId,
		Payload: psutil.Detail(true),
	})

	return err

}
