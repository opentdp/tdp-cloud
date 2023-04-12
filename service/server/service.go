package server

import (
	"github.com/kardianos/service"
	"github.com/open-tdp/go-helper/logman"

	"tdp-cloud/cmd/args"
)

var svclog service.Logger

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
		logman.Fatal("Init service failed", "error", err)
	}

	svclog, err = svc.Logger(nil)
	if err != nil {
		logman.Fatal("Init service failed", "error", err)
	}

	return svc

}
