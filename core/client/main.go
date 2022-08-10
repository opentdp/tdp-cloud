package client

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type H map[string]any

func Connect(url string) {

	log.Println("客户端模式暂未实现")

	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		var v = H{"e": "error"}
		for {
			err := ws.WriteJSON(v)
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		var v = H{}
		err := ws.ReadJSON(&v)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("receive: ", v)
	}

}
