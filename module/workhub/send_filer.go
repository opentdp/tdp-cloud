package workhub

import (
	"fmt"
	"time"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/socket"

	"tdp-cloud/module/fsadmin"
)

func (pod *SendPod) Filer(data *fsadmin.FilerParam) (string, error) {

	var (
		err    error
		taskId = uint(time.Now().UnixNano())
	)

	logman.Info("filer:send", "to", pod.WorkerMeta.HostName)

	err = pod.WriteJson(&socket.PlainData{
		Method:  "Filer",
		TaskId:  taskId,
		Payload: data,
	})

	id := "filer" + fmt.Sprintf("%d", taskId)
	return id, err

}

func (pod *RespPod) Filer(rs *socket.PlainData) {

	logman.Info("filer:resp", "from", pod.WorkerMeta.HostName)

	id := "filer" + fmt.Sprintf("%d", rs.TaskId)
	workerResp[id] = rs

}
