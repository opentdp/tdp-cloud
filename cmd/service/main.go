package service

import (
	"log"
)

func Create() {

	if len(vInstall) > 0 {
		logPrint("install", serviceInstall())
		return
	}

	if len(vUninstall) > 0 {
		logPrint("uninstall", serviceUninstall())
		return
	}

}

func serviceInstall() error {

	switch vInstall {
	case "server":
		return serverService().Install()
	case "worker":
		return workerService().Install()
	}

	return nil

}

func serviceUninstall() error {

	switch vUninstall {
	case "server":
		return serverService().Uninstall()
	case "worker":
		return workerService().Uninstall()
	}

	return nil

}

func logPrint(n string, e error) {

	if e != nil {
		log.Print(n, "service error:", e.Error())
	} else {
		log.Print(n, "service done")
	}

}
