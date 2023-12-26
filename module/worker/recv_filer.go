package worker

import (
	"github.com/opentdp/go-helper/filer"
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/socket"

	"tdp-cloud/module/fsadmin"
)

func (pod *RecvPod) Filer(rq *socket.PlainData) error {

	var (
		err  error
		msg  string
		ret  []*filer.FileInfo
		data fsadmin.FilerParam
	)

	logman.Info("filer:recv", "taskId", rq.TaskId)

	if err = rq.GetPayload(&data); err == nil {
		ret, err = fsadmin.Filer(&data)
	}

	if err != nil {
		msg = err.Error()
	}

	err = pod.WriteJson(&socket.PlainData{
		Method:  "Filer:resp",
		TaskId:  rq.TaskId,
		Success: err == nil,
		Message: msg,
		Payload: ret,
	})

	return err

}
