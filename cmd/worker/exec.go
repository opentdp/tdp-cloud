package worker

import (
	"log"
	"time"

	"github.com/spf13/viper"

	"tdp-cloud/module/worker"
)

func Execute() {

	defer timer()

	remote := viper.GetString("worker.remote")

	if err := worker.Daemon(remote); err != nil {
		log.Println(err)
	}

}

func timer() {

	log.Println("连接已断开，将在5秒后重试")
	time.Sleep(time.Second * 5)
	Execute()

}
