package worker

import (
	"log"
	"time"

	"tdp-cloud/internal/worker"
)

func Connect(url string) {

	defer delayer(url)

	worker.Register(url)

}

func delayer(url string) {

	log.Println("连接失败，将在5秒后重试")

	time.Sleep(time.Second * 5)
	Connect(url)

}
