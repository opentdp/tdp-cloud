package server

import (
	"log"

	"github.com/kardianos/service"
	"github.com/spf13/viper"
)

func Service(args []string) service.Service {

	config := &service.Config{
		Name:        "tdp-server",
		DisplayName: "TDP Cloud Server",
		Description: "TDP Control Panel Server",
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
