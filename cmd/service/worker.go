package service

import (
	"log"
	"os"

	"github.com/kardianos/service"
)

type worker struct{}

func (p *worker) Start(s service.Service) error {
	return nil
}

func (p *worker) Stop(s service.Service) error {
	return nil
}

func workerService() service.Service {

	var args = []string{os.Args[3]}

	if len(os.Args) > 5 {
		args = append(args, os.Args[5:]...)
	}

	svcConfig := &service.Config{
		Name:        "tdp-cloud",
		DisplayName: "tdp cloud worker",
		Description: "tdp cloud worker",
		Arguments:   args,
	}

	s, err := service.New(&worker{}, svcConfig)

	if err != nil {
		log.Fatal("init service error:", err)
	}

	return s

}

func workerInstall() {

	s := workerService()

	if x := s.Install(); x != nil {
		log.Print("install service error:", x.Error())
	} else {
		log.Print("install service done")
	}

}

func workerUninstall() {

	s := workerService()

	if x := s.Uninstall(); x != nil {
		log.Print("uninstall service error:", x.Error())
	} else {
		log.Print("uninstall service done")
	}

}
