package worker

import (
	"log"
	"time"

	"tdp-cloud/cmd/args"

	"tdp-cloud/internal/worker"
)

func Create() {

	defer delayer()

	worker.Register(args.Server)

}

func delayer() {

	log.Println("连接失败，将在5秒后重试")

	time.Sleep(time.Second * 5)
	Create()

}
