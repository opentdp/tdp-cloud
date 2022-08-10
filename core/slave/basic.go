package slave

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

type SocketData struct {
	Action  string
	Method  string
	Payload any
}

func Connect(url string) {

	log.Println("客户端模式暂未实现，仅供调试使用")

	// 断线重连
	defer delayConnect(url)

	// 连接服务器
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return
	}

	// 延迟关闭连接
	defer ws.Close()

	// 保持连接

	go func() {
		for {
			if err := Ping(ws); err != nil {
				log.Println(err)
				break
			}
			time.Sleep(time.Second * 30)
		}
	}()

	// 接收数据

	for {
		var rq SocketData

		if ws.ReadJSON(&rq) != nil {
			break
		}

		switch rq.Action {
		case "runCommand":
			var data *CommandPayload
			if mapstructure.Decode(rq.Payload, &data) == nil {
				RunCommand(ws, data)
			} else {
				log.Println("runCommand 参数错误")
			}
		default:
			log.Println("receive: ", rq)
		}
	}

}

func delayConnect(url string) {

	log.Println("连接失败，将在5秒后重试")

	time.Sleep(time.Second * 5)
	Connect(url)

}
