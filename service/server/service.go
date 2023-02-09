package server

import (
	"log"
	"os"

	"github.com/kardianos/service"
)

func Service() service.Service {

	var args = []string{"server"}

	if len(os.Args) > 4 {
		args = append(args, os.Args[4:]...)
	}

	config := &service.Config{
		Name:        "tdp-server",
		DisplayName: "TDP Cloud Server",
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
