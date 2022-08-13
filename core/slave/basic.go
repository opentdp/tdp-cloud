package slave

import (
	"log"
	"time"

	"tdp-cloud/core/serve/agent"
	"tdp-cloud/core/socket"
)

type RecvPod agent.RecvPod
type SendPod agent.SendPod

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

	send := NewSendPod(pod)

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

	recv := NewRecvPod(pod)

	for {
		var rq *SocketData

		if pod.Read(&rq) != nil {
			break
		}

		switch rq.Method {
		case "RunCommand":
			recv.RunCommand(rq)
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

func NewRecvPod(pod *socket.JsonPod) *RecvPod {

	return &RecvPod{pod}

}

func NewSendPod(pod *socket.JsonPod) *SendPod {

	return &SendPod{pod}

}
