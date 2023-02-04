package installer

import (
	"log"
)

func Execute(iName, uName string) {

	if len(uName) > 0 {
		logPrint("uninstall", uninstall(uName))
	}

	if len(iName) > 0 {
		logPrint("install", install(iName))
	}

}

func install(name string) error {

	switch name {
	case "server":
		return serverService().Install()
	case "worker":
		return workerService().Install()
	}

	return nil

}

func uninstall(name string) error {

	switch name {
	case "server":
		return serverService().Uninstall()
	case "worker":
		return workerService().Uninstall()
	}

	return nil

}

func logPrint(n string, e error) {

	if e != nil {
		log.Fatalln(n, "service error:", e.Error())
	} else {
		log.Println(n, "service done")
	}

}
