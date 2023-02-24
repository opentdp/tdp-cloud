package worker

import (
	"time"

	"github.com/kardianos/service"
	"github.com/spf13/viper"

	"tdp-cloud/helper/logman"
	"tdp-cloud/module/worker"
)

type program struct{}

func (p *program) Start(s service.Service) error {

	logman.Info("TDP Worker start")

	go p.run()
	return nil

}

func (p *program) Stop(s service.Service) error {

	logman.Info("TDP Worker stop")

	return nil

}

func (p *program) run() {

	defer p.timer()

	remote := viper.GetString("worker.remote")

	if err := worker.Daemon(remote); err != nil {
		logman.Info(err)
	}

}

func (p *program) timer() {

	logman.Info("Connection disconnected, retry in 5 seconds.")

	time.Sleep(5 * time.Second)
	p.run()

}
