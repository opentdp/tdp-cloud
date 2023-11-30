package workhub

import (
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/psutil"
	"github.com/opentdp/go-helper/socket"
)

func (pod *RecvPod) Ping(rq *socket.PlainData) error {

	var (
		err  error
		stat = psutil.SummaryStat{}
	)

	logman.Info("ping:recv", "from", pod.Conn.RemoteAddr())

	if rq.GetPayload(&stat) == nil {
		pod.WorkerMeta = &stat
	}

	err = pod.WriteJson(&socket.PlainData{
		Method:  "Ping:resp",
		TaskId:  rq.TaskId,
		Success: true,
		Payload: "OK",
	})

	return err

}
