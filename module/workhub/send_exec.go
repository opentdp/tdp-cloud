package workhub

import (
	"github.com/opentdp/go-helper/command"
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/socket"
)

func (pod *SendPod) Exec(data *command.ExecPayload) (uint, error) {

	var (
		err    error
		taskId = createHistory(pod, data)
	)

	logman.Info("exec:send", "to", pod.WorkerMeta.HostName)

	err = pod.WriteJson(&socket.PlainData{
		Method:  "Exec",
		TaskId:  taskId,
		Payload: data,
	})

	return taskId, err

}

func (pod *RespPod) Exec(rs *socket.PlainData) {

	logman.Info("exec:resp", "from", pod.WorkerMeta.HostName)

	err := updateHistory(pod, rs)

	logman.Info("exec:save", "err", err)

}
