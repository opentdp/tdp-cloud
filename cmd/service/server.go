package service

import (
	"log"
	"os"

	"github.com/kardianos/service"
)

type server struct{}

func (p *server) Start(s service.Service) error {

	log.Print("service start")
	return nil

}

func (p *server) Stop(s service.Service) error {

	log.Print("service stop")
	return nil

}

func serverService() service.Service {

	var args = []string{os.Args[3]}

	if len(os.Args) > 5 {
		args = append(args, os.Args[5:]...)
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

	s, err := service.New(&server{}, config)

	if err != nil {
		log.Fatal("init service error: ", err)
	}

	return s

}
