package worker

import (
	"log"

	"tdp-cloud/helper/psutil"
)

func (pod *RecvPod) Stat(rs *SocketData) error {

	rq := &SocketData{
		Method:  "Stat:resp",
		TaskId:  rs.TaskId,
		Payload: psutil.Detail(),
	}

	if err := pod.Write(rq); err != nil {
		log.Println("Exec:resp", err)
		return err
	}

	return nil

}
