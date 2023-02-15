package worker

import (
	"log"
	"time"

	"tdp-cloud/helper/psutil"
)

func (pod *SendPod) Ping() (uint, error) {

	rq := &SocketData{
		Method:  "Ping",
		TaskId:  0,
		Payload: psutil.Summary(),
	}

	log.Println("Ping:send", "SummaryStat")

	return rq.TaskId, pod.Write(rq)

}

func (pod *RespPod) Ping(rs *SocketData) {

	log.Println("Ping:resp", rs.Payload)

}

//// 持续报送状态

func PingLoop(pod *SendPod) error {

	for {
		if _, err := pod.Ping(); err != nil {
			log.Println("Ping:error", err)
			return err
		}
		time.Sleep(25 * time.Second)
	}

}
