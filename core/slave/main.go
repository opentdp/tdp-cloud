package slave

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func Connect(url string) {

	log.Println("客户端模式暂未实现")

	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Println("无法连接，将在5秒后重试")
		time.Sleep(time.Second * 5)
		Connect(url)
		return
	}

	// 断线重连

	defer func() {
		log.Println("连接断开，将在5秒后重试")
		time.Sleep(time.Second * 5)
		Connect(url)
	}()

	// 保持连接

	go func() {
		var v = SocketData{
			Action:  "ping",
			Method:  "request",
			Payload: "xx",
		}
		for {
			log.Println("send: ", v)
			if err := ws.WriteJSON(v); err != nil {
				log.Println(err)
				break
			}
			time.Sleep(time.Second * 30)
		}
	}()

	// 接收数据

	for {
		var v = SocketData{}
		if err := ws.ReadJSON(&v); err != nil {
			log.Println(err)
			break
		}
		log.Println("receive: ", v)
	}

}
