package worker

import (
	"log"

	"tdp-cloud/helper/psutil"
)

func (pod *RecvPod) Stat(rs *SocketData) error {

	log.Println("Stat:recv Id", rs.TaskId)

	rq := &SocketData{
		Method:  "Stat:resp",
		TaskId:  rs.TaskId,
		Payload: psutil.Detail(),
	}

	return pod.Write(rq)

}
