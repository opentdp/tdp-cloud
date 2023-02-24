package worker

import (
	"github.com/kardianos/service"
	"github.com/spf13/viper"

	"tdp-cloud/helper/logman"
)

func Service(args []string) service.Service {

	config := &service.Config{
		Name:        "tdp-worker",
		DisplayName: "TDP Cloud Worker",
		Description: "TDP Control Panel Worker",
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
