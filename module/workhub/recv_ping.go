package workhub

import (
	"github.com/mitchellh/mapstructure"
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/psutil"
)

func (pod *RecvPod) Ping(rq *SocketData) error {

	logman.Info("ping:recv", "from", pod.Conn.RemoteAddr())

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
