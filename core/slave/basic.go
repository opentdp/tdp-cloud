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

	log.Println("客户端模式暂未实现，仅供调试使用")

	// 自动重连
	defer delayConnect(url)

	// 注册服务

	pod, err := socket.NewJsonPodClient(url)

	if err != nil {
		return
	}

	defer pod.Close()

	// 保持连接

	send := &SendPod{pod}

	go func() {
		for {
			log.Println("send: Ping")
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

func delayConnect(url string) {

	log.Println("连接失败，将在5秒后重试")

	time.Sleep(time.Second * 5)
	Connect(url)

}
