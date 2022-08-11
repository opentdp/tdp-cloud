package slave

import (
	"log"
	"time"

	"github.com/mitchellh/mapstructure"

	"tdp-cloud/core/socket"
)

type SocketData struct {
	Action  string
	Method  string
	Payload any
	Error   error
}

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

	go func() {
		for {
			log.Println("send: Ping")
			if err := Ping(pod); err != nil {
				log.Println(err)
				break
			}
			time.Sleep(time.Second * 30)
		}
	}()

	// 接收数据

	for {
		var rq *SocketData

		if pod.Read(&rq) != nil {
			break
		}

		switch rq.Action {
		case "runCommand":
			var data *CommandPayload
			if mapstructure.Decode(rq.Payload, &data) == nil {
				RunCommand(pod, data)
			} else {
				log.Println("runCommand 参数错误")
			}
		case "pong":
			log.Println("receive: Pong - ", rq.Payload)
		default:
			log.Println("unkown: ", rq)
		}
	}

}

func delayConnect(url string) {

	log.Println("连接失败，将在5秒后重试")

	time.Sleep(time.Second * 5)
	Connect(url)

}
