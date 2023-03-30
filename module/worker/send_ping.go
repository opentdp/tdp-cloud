package worker

import (
	"time"

	"tdp-cloud/helper/logman"
	"tdp-cloud/helper/psutil"
)

func (pod *SendPod) Ping() (uint, error) {

	logman.Info("Ping:send", "SummaryStat")

	info := psutil.Summary(true)

	rq := &SocketData{
		Method:  "Ping",
		TaskId:  0,
		Payload: info,
	}

	return rq.TaskId, pod.Write(rq)

}

func (pod *RespPod) Ping(rs *SocketData) {

	logman.Info("Ping:resp", rs.Payload)

}

//// 持续报送状态

func PingLoop(pod *SendPod) error {

	for {
		if _, err := pod.Ping(); err != nil {
			logman.Error("Ping:fail", err)
			return err
		}
		time.Sleep(25 * time.Second)
	}

}
