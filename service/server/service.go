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
		Option: service.KeyValue{
			"LogDirectory": viper.GetString("logger.dir"),
			"LogOutput":    viper.GetBool("logger.output"),
		},
		Arguments: args,
	}

	svc, err := service.New(&program{}, config)

	if err != nil {
		log.Fatalln("Init service error:", err)
	}

	return svc

}
