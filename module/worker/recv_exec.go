package worker

import (
	"errors"
	"log"

	"github.com/mitchellh/mapstructure"

	"tdp-cloud/helper/command"
)

func (pod *RecvPod) Exec(rs *SocketData) error {

	var (
		err  error
		ret  string
		data *command.ExecPayload
	)

	log.Println("Exec:recv", rs.Payload)

	if mapstructure.Decode(rs.Payload, &data) == nil {
		ret, err = command.Exec(data)
	} else {
		err = errors.New("无法解析请求参数")
	}

	if err != nil {
		log.Println("Exec:fail", err)
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

	return pod.Write(rq)
}
