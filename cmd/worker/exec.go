package worker

import (
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"tdp-cloud/module/service"
	"tdp-cloud/module/worker"
)

func Execute(cmd *cobra.Command, params []string) {

	switch svc {
	case "install":
		service.Worker().Install()
	case "uninstall":
		service.Worker().Uninstall()
	case "":
		start()
	}

}

func start() {

	defer timer()

	remote := viper.GetString("worker.remote")

	if err := worker.Daemon(remote); err != nil {
		log.Println(err)
	}

}

func timer() {

	log.Println("连接已断开，将在5秒后重试")
	time.Sleep(time.Second * 5)
	start()

}
