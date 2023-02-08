package worker

import (
	"errors"
	"log"

	"github.com/mitchellh/mapstructure"

	"tdp-cloud/helper/command"
)

func (pod *RecvPod) Exec(rs *SocketData) error {

	var err error
	var ret string

	var data *command.ExecPayload

	if mapstructure.Decode(rs.Payload, &data) == nil {
		log.Println("Exec:wait", data.Name)
		ret, err = command.Exec(data)
	} else {
		err = errors.New("无法解析请求参数")
	}

	if err != nil {
		log.Println("Exec:error", err)
	} else {
		log.Println("Exec:done", data.Name)
	}

	rq := &SocketData{
		Method:  "Exec:resp",
		TaskId:  rs.TaskId,
		Success: err == nil,
		Payload: map[string]any{
			"Output": ret,
			"Error":  err,
		},
	}

	if err = pod.Write(rq); err != nil {
		log.Println("Exec:resp", err)
	}

	return err

}
