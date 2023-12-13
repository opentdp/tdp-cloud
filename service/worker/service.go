package worker

import (
	"github.com/kardianos/service"
	"github.com/opentdp/go-helper/logman"

	"tdp-cloud/cmd/args"
)

func Service(param []string) service.Service {

	config := &service.Config{
		Name:        "tdp-worker",
		DisplayName: "TDP Cloud Worker",
		Description: "TDP Control Panel Worker",
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
		logger.Info("TDP Worker start")
	} else {
		logman.Info("TDP Worker start")
	}

	go origin()
	return nil

}

func (p *program) Stop(s service.Service) error {

	if logger, err := s.Logger(nil); err == nil {
		logger.Info("TDP Worker stop")
	} else {
		logman.Info("TDP Worker stop")
	}

	return nil

}
