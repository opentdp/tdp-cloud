package server

import (
	"github.com/kardianos/service"
	"github.com/spf13/viper"

	"tdp-cloud/helper/logman"
)

func Service(args []string) service.Service {

	config := &service.Config{
		Name:        "tdp-server",
		DisplayName: "TDP Cloud Server",
		Description: "TDP Control Panel Server",
		Option: service.KeyValue{
			"LogDirectory": viper.GetString("logger.dir"),
			"LogOutput":    viper.GetBool("logger.tofile"),
		},
		Arguments: args,
	}

	svc, err := service.New(&program{}, config)

	if err != nil {
		logman.Fatal("Init service error:", err)
	}

	return svc

}
