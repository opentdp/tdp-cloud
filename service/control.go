package service

import (
	"os"

	"github.com/kardianos/service"

	"tdp-cloud/cmd/args"
	"tdp-cloud/helper/logman"
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

	// 获取抽象类

	switch name {
	case "server":
		svc = server.Service(cliArgs())
	case "worker":
		svc = worker.Service(cliArgs())
	default:
		logman.Fatal("Unknown service", "Name", name)
	}

	// 强制保存配置

	if act == "" || act == "start" {
		args.MustSave()
	}

	// 执行服务动作

	switch act {
	case "": // 直接运行
		if err := svc.Run(); err != nil {
			logman.Fatal(svc.String()+" run failed", "Error", err)
		}
	case "status": // 查看状态
		if sta, err := svc.Status(); err != nil {
			logman.Fatal(svc.String()+" "+act+" failed", "Error", err)
		} else {
			logman.Warn(svc.String(), "Status", statusMap[sta])
		}
	default: // 其他动作
		if err := service.Control(svc, act); err != nil {
			logman.Fatal(svc.String()+" "+act+" failed", "Error", err)
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
