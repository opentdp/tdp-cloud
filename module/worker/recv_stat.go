package worker

import (
	"runtime"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/psutil"
	"github.com/opentdp/go-helper/socket"
)

func (pod *RecvPod) Stat(rq *socket.PlainData) error {

	var err error

	logman.Info("stat:recv", "taskId", rq.TaskId)

	err = pod.WriteJson(&socket.PlainData{
		Method:  "Stat:resp",
		TaskId:  rq.TaskId,
		Success: true,
		Payload: map[string]any{
			"Stat":         psutil.Detail(true),
			"MemStat":      psutil.GoMemory(),
			"NumGoroutine": runtime.NumGoroutine(),
		},
	})

	return err

}
