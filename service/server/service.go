package server

import (
	"github.com/kardianos/service"
	"github.com/opentdp/go-helper/logman"

	"tdp-cloud/cmd/args"
)

func Service(param []string) service.Service {

	config := &service.Config{
		Name:        "tdp-server",
		DisplayName: "TDP Cloud Server",
		Description: "TDP Control Panel Server",
		Option: service.KeyValue{
			"LogDirectory": args.Logger.Dir,
			"LogOutput":    args.Logger.Target == "file",
		},
		Arguments: param,
	}

	svc, err := service.New(&program{}, config)
	if err != nil {
		logman.Fatal("init service failed", "error", err)
	}

	return svc

}

// service program

type program struct{}

func (p *program) Start(s service.Service) error {

	if logger, err := s.Logger(nil); err == nil {
		logger.Info("TDP Server start")
	} else {
		logman.Info("TDP Server start")
	}

	go origin()
	return nil

}

func (p *program) Stop(s service.Service) error {

	if logger, err := s.Logger(nil); err == nil {
		logger.Info("TDP Server stop")
	} else {
		logman.Info("TDP Server stop")
	}

	return nil

}
