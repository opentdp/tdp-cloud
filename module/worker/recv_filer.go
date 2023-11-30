package worker

import (
	"errors"
	"os"

	"github.com/opentdp/go-helper/filer"
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/socket"

	"tdp-cloud/module/workhub"
)

func (pod *RecvPod) Filer(rq *socket.PlainData) error {

	var (
		err error
		msg string
		ret struct {
			Success  bool
			FileData []byte
			FileList []*filer.FileInfo
		}
		data workhub.FilerPayload
	)

	logman.Info("filer:recv", "taskId", rq.TaskId)

	if err = rq.GetPayload(&data); err == nil {
		switch data.Action {
		case "ls":
			ret.FileList, err = filer.List(data.Path)
		case "read":
			ret.FileData, err = os.ReadFile(data.Path)
		case "write":
			err = filer.Write(data.Path, data.File.Data)
		case "chmod":
			err = os.Chmod(data.Path, data.File.Mode)
		case "mkdir":
			err = os.MkdirAll(data.Path, 0755)
		case "rm":
			err = os.RemoveAll(data.Path)
		case "mv":
			err = os.Rename(data.Path, data.File.Name)
		default:
			err = errors.New("无法识别的操作")
		}
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
