package worker

import (
	"log"
	"time"

	"tdp-cloud/module/worker"
)

func Create() {

	defer delayer()

	if err := worker.Daemon(vRemote); err != nil {
		log.Print(err)
	}

}

func delayer() {

	log.Println("连接已断开，将在5秒后重试")

	time.Sleep(time.Second * 5)
	Create()

}
