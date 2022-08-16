package slave

import (
	"log"
	"time"

	"tdp-cloud/core/socket"
)

func Connect(url string) {

	defer delayer(url)

	pod, err := socket.NewJsonPodClient(url)

	if err != nil {
		return
	}

	defer pod.Close()

	go Sender(pod)

	Receiver(pod)

}

func delayer(url string) {

	log.Println("连接失败，将在5秒后重试")

	time.Sleep(time.Second * 5)
	Connect(url)

}
