package service

import (
	"log"
)

func Create() {

	if len(vInstall) > 0 {
		logPrint("install", install())
		return
	}

	if len(vUninstall) > 0 {
		logPrint("uninstall", uninstall())
		return
	}

}

func install() error {

	switch vInstall {
	case "server":
		return serverService().Install()
	case "worker":
		return workerService().Install()
	}

	return nil

}

func uninstall() error {

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
		log.Print(n, " service error: ", e.Error())
	} else {
		log.Print(n, " service done")
	}

}
