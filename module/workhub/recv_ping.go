package workhub

import (
	"github.com/mitchellh/mapstructure"

	"tdp-cloud/helper/logman"
	"tdp-cloud/helper/psutil"
)

func (pod *RecvPod) Ping(rq *SocketData) error {

	logman.Info("Ping:recv", "From", pod.Conn.RemoteAddr())

	stat := &psutil.SummaryStat{}
	if mapstructure.Decode(rq.Payload, stat) == nil {
		pod.WorkerMeta = stat
	}

	err := pod.WriteJson(&SocketData{
		Method:  "Ping:resp",
		TaskId:  rq.TaskId,
		Payload: "OK",
	})

	return err

}
