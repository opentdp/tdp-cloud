package worker

import (
	"github.com/kardianos/service"

	"tdp-cloud/cmd/args"
	"tdp-cloud/helper/logman"
)

var svclog service.Logger

func Service(param []string) service.Service {

	config := &service.Config{
		Name:        "tdp-worker",
		DisplayName: "TDP Cloud Worker",
		Description: "TDP Control Panel Worker",
		Option: service.KeyValue{
			"LogDirectory": args.Logger.Dir,
			"LogOutput":    args.Logger.ToFile,
		},
		Arguments: param,
	}

	svc, err := service.New(&program{}, config)
	if err != nil {
		logman.Fatal("Init service error:", err)
	}

	svclog, err = svc.Logger(nil)
	if err != nil {
		logman.Fatal("Init service error:", err)
	}

	return svc

}
