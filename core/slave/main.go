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
		log.Fatal(err)
	}

	go func() {
		var v = SocketData{
			Action:  "ping",
			Method:  "request",
			Payload: "xx",
		}
		for {
			log.Println("send: ", v)
			if err := ws.WriteJSON(v); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second * 30)
		}
	}()

	for {
		var v = SocketData{}
		if err := ws.ReadJSON(&v); err != nil {
			log.Fatal(err)
		}
		log.Println("receive: ", v)
	}

}
