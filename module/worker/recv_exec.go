package worker

import (
	"errors"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/opentdp/go-helper/command"
	"github.com/opentdp/go-helper/logman"
)

func (pod *RecvPod) Exec(rs *SocketData) error {

	var (
		err  error
		ret  string
		data *command.ExecPayload
	)

	logman.Info("exec:recv", "payload", rs.Payload)

	if mapstructure.Decode(rs.Payload, &data) == nil {
		ret, err = command.Exec(data)
	} else {
		err = errors.New("无法解析请求参数")
	}

	if err != nil {
		logman.Error("exec:fail", "error", err)
	} else {
		logman.Info("exec:done", "name", data.Name)
	}

	err = pod.WriteJson(&SocketData{
		Method:  "Exec:resp",
		TaskId:  rs.TaskId,
		Message: err,
		Payload: strings.TrimSpace(ret),
	})

	return err

}
