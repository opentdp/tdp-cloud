package service

import (
	"log"
	"os"

	"github.com/kardianos/service"
)

type worker struct{}

func (p *worker) Start(s service.Service) error {

	log.Print("service start")
	return nil

}

func (p *worker) Stop(s service.Service) error {

	log.Print("service stop")
	return nil

}

func workerService() service.Service {

	var args = []string{os.Args[3]}

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

	s, err := service.New(&worker{}, config)

	if err != nil {
		log.Fatal("init service error: ", err)
	}

	return s

}
