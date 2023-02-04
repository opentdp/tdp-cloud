package service

import (
	"log"
	"os"

	"github.com/kardianos/service"
)

type server struct{}

func (p *server) Start(s service.Service) error {

	log.Println("service start")
	return nil

}

func (p *server) Stop(s service.Service) error {

	log.Println("service stop")
	return nil

}

func Server() service.Service {

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
		log.Fatalln("Init service error:", err)
	}

	return s

}
