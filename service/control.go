package service

import (
	"log"

	"github.com/kardianos/service"

	"tdp-cloud/helper/strutil"
	"tdp-cloud/service/server"
	"tdp-cloud/service/worker"
)

var statusMap = map[service.Status]string{
	0: "Unknown",
	1: "Running",
	2: "Stopped",
}

func Control(name, act string) {

	var svc service.Service

	switch name {
	case "server":
		svc = server.Service()
	case "worker":
		svc = worker.Service()
	default:
		log.Fatalln("未知服务")
	}

	log.Println(strutil.FirstUpper(act), "service", svc.String(), "...")

	switch act {
	case "":
		if err := svc.Run(); err != nil {
			log.Fatalln(err)
		}
	case "status":
		if sta, err := svc.Status(); err == nil {
			log.Println(svc.String(), "Status:", statusMap[sta])
		} else {
			log.Fatalln(err)
		}
	default:
		if err := service.Control(svc, act); err != nil {
			log.Fatalln(err)
		}
	}

}
