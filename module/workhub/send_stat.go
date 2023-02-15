package workhub

import (
	"log"
)

func (pod *SendPod) Stat() (uint, error) {

	rq := &SocketData{
		Method: "Stat",
		TaskId: 0,
	}

	return rq.TaskId, pod.Write(rq)

}

func (pod *RespPod) Stat(rq *SocketData) {

	log.Println(rq)

}
