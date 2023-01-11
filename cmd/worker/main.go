package worker

import (
	"log"
	"time"

	"tdp-cloud/internal/worker"
)

func Create() {

	defer delayer()

	worker.Daemon(vRemote)

}

func delayer() {

	log.Println("连接失败，将在5秒后重试")

	time.Sleep(time.Second * 5)
	Create()

}
