package worker

import (
	"log"
	"os"

	"github.com/kardianos/service"
)

func Service() service.Service {

	var args = []string{"worker"}

	if len(os.Args) > 5 {
		args = append(args, os.Args[5:]...)
	}

	config := &service.Config{
		Name:        "tdp-worker",
		DisplayName: "TDP Cloud Worker",
		Description: "TDP Cloud Control Panel",
		Arguments:   args,
		Option: service.KeyValue{
			"LogDirectory": "/var/log/tdp-cloud",
		},
	}

	svc, err := service.New(&origin{}, config)

	if err != nil {
		log.Fatalln("Init service error:", err)
	}

	return svc

}
