package worker

import (
	"log"

	"github.com/kardianos/service"
	"github.com/spf13/viper"
)

func Service(args []string) service.Service {

	config := &service.Config{
		Name:        "tdp-worker",
		DisplayName: "TDP Cloud Worker",
		Description: "TDP Control Panel Worker",
		Option:      service.KeyValue{},
		Arguments:   args,
	}

	if logPath := viper.GetString("logger.directory"); logPath != "" {
		config.Option["LogDirectory"] = logPath
	}

	svc, err := service.New(&program{}, config)

	if err != nil {
		log.Fatalln("Init service error:", err)
	}

	return svc

}
