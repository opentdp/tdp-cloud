package service

import (
	"log"
	"os"

	"github.com/kardianos/service"

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

	// 获取服务抽象类

	switch name {
	case "server":
		svc = server.Service(cliArgs())
	case "worker":
		svc = worker.Service(cliArgs())
	default:
		log.Fatalln("Unknown service:", name)
	}

	// 执行服务动作

	switch act {
	case "": // 直接运行
		if err := svc.Run(); err != nil {
			log.Fatalln(err)
		}
	case "status": // 查看状态
		if sta, err := svc.Status(); err == nil {
			log.Println(svc.String(), "Status:", statusMap[sta])
		} else {
			log.Fatalln(err)
		}
	default: // 其他动作
		if err := service.Control(svc, act); err != nil {
			log.Fatalln(err)
		}
	}

}

func cliArgs() []string {

	args := []string{}

	for i, n := 1, len(os.Args); i < n; i++ {
		if v := os.Args[i]; v != "-s" && v != "--service" {
			args = append(args, v)
		} else {
			i++
		}
	}

	return args

}
