package server

import (
	"github.com/kardianos/service"

	"tdp-cloud/cmd/args"
	"tdp-cloud/helper/logman"
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
		logman.Fatal("Init service failed", "Error", err)
	}

	svclog, err = svc.Logger(nil)
	if err != nil {
		logman.Fatal("Init service failed", "Error", err)
	}

	return svc

}
