package worker

import (
	"strings"

	"github.com/opentdp/go-helper/command"
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/socket"
)

func (pod *RecvPod) Exec(rq *socket.PlainData) error {

	var (
		err  error
		msg  string
		ret  string
		data command.ExecPayload
	)

	logman.Info("exec:recv", "payload", rq.Payload)

	if err = rq.GetPayload(&data); err == nil {
		ret, err = command.Exec(&data)
	}

	if err != nil {
		msg = err.Error()
		logman.Error("exec:fail", "error", err)
	} else {
		logman.Info("exec:done", "name", data.Name)
	}

	ret = strings.TrimSpace(ret)
	err = pod.WriteJson(&socket.PlainData{
		Method:  "Exec:resp",
		TaskId:  rq.TaskId,
		Success: err == nil,
		Message: msg,
		Payload: ret,
	})

	return err

}
