package worker

import (
	"errors"

	"github.com/mitchellh/mapstructure"
	"github.com/opentdp/go-helper/filer"
	"github.com/opentdp/go-helper/logman"

	"tdp-cloud/module/workhub"
)

func (pod *RecvPod) Filer(rs *SocketData) error {

	var (
		err  error
		ret  []*filer.FileInfo
		data workhub.FilerPayload
	)

	logman.Info("filer:recv", "taskId", rs.TaskId)

	if mapstructure.Decode(rs.Payload, &data) == nil {
		ret, err = filer.List(data.Path)
	} else {
		err = errors.New("无法解析请求参数")
	}

	err = pod.WriteJson(&SocketData{
		Method:  "Filer:resp",
		TaskId:  rs.TaskId,
		Message: err,
		Payload: ret,
	})

	return err

}
