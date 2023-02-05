package subset

import (
	"log"
	"tdp-cloud/helper/strutil"

	"github.com/kardianos/service"
)

var statusMap = map[service.Status]string{
	0: "Unknown",
	1: "Running",
	2: "Stopped",
}

func ctrl(svc service.Service, act string) {

	log.Println(strutil.FirstUpper(act), "service", svc.String(), "...")

	switch act {
	case "":
		err := svc.Run()
		if err != nil {
			log.Fatalln(err)
		}
	case "status":
		sta, err := svc.Status()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(svc.String(), "Status:", statusMap[sta])
	default:
		err := service.Control(svc, act)
		if err != nil {
			log.Fatalln(err)
		}
	}

}
