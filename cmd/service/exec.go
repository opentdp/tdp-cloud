package service

import (
	"log"

	"github.com/spf13/viper"

	"tdp-cloud/module/service"
)

func Execute() {

	if in := viper.GetString("install"); in != "" {
		logPrint("uninstall", uninstall(in))
	}

	if un := viper.GetString("uninstall"); un != "" {
		logPrint("install", install(un))
	}

}

func logPrint(n string, e error) {

	if e != nil {
		log.Fatalln(n, "service error:", e.Error())
	} else {
		log.Println(n, "service done")
	}

}

func install(name string) error {

	switch name {
	case "server":
		return service.Server().Install()
	case "worker":
		return service.Worker().Install()
	}

	return nil

}

func uninstall(name string) error {

	switch name {
	case "server":
		return service.Server().Uninstall()
	case "worker":
		return service.Worker().Uninstall()
	}

	return nil

}
