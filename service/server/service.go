package server

import (
	"log"

	"github.com/kardianos/service"
)

func Service(args []string) service.Service {

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
