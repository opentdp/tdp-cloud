package worker

import (
	"log"
	"time"

	"tdp-cloud/helper/psutil"
)

func (pod *SendPod) Ping() (uint, error) {

	log.Println("Ping:send", "SummaryStat")

	rq := &SocketData{
		Method:  "Ping",
		TaskId:  0,
		Payload: psutil.Summary(),
	}

	return rq.TaskId, pod.Write(rq)

}

func (pod *RespPod) Ping(rs *SocketData) {

	log.Println("Ping:resp", rs.Payload)

}

//// 持续报送状态

func PingLoop(pod *SendPod) error {

	for {
		if _, err := pod.Ping(); err != nil {
			log.Println("Ping:fail", err)
			return err
		}
		time.Sleep(25 * time.Second)
	}

}
