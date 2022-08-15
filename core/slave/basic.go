package slave

import (
	"log"
	"time"

	"tdp-cloud/core/serve/agent"
	"tdp-cloud/core/socket"
)

type RecvPod struct {
	*socket.JsonPod
}

type RespPod struct {
	*socket.JsonPod
}

type SendPod struct {
	*socket.JsonPod
}

type SocketData agent.SocketData

func Connect(url string) {

	defer delayer(url)

	// 注册服务

	pod, err := socket.NewJsonPodClient(url)

	if err != nil {
		return
	}

	defer pod.Close()

	// 发送数据

	send := &SendPod{pod}

	if _, err := send.Register(); err != nil {
		return
	}

	go func() {
		for {
			if _, err := send.Ping(); err != nil {
				log.Println(err)
				break
			}
			time.Sleep(time.Second * 15)
		}
	}()

	// 接收数据

	recv := &RecvPod{pod}
	resp := &RespPod{pod}

	for {
		var rq *SocketData

		if pod.Read(&rq) != nil {
			break
		}

		switch rq.Method {
		case "Exec":
			recv.Exec(rq)
		case "Ping:resp":
			resp.Ping(rq)
		default:
			log.Println("recv:", rq)
		}
	}

}

func delayer(url string) {

	log.Println("连接失败，将在5秒后重试")

	time.Sleep(time.Second * 5)
	Connect(url)

}
